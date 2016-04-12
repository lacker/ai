#!/usr/bin/env python3
import random
import tensorflow as tf
import numpy as np

from tensorflow.models.rnn import rnn_cell
from tensorflow.models.rnn import seq2seq

MIN_NUMBER = 1
MAX_NUMBER = 127
def number():
    return random.randrange(MIN_NUMBER, MAX_NUMBER + 1)

# Format is like
# 10*10
# 100
# space-padded on the right to get consistent lengths.
SOURCE_VOCAB, TARGET_VOCAB = '10* ', '10 '
SOURCE_LEN = 1 + 2 * len('{0:b}'.format(MAX_NUMBER))
TARGET_LEN = len('{0:b}'.format(MAX_NUMBER * MAX_NUMBER))

def source_pad(s):
  while len(s) <= SOURCE_LEN:
    s += ' '
  return s

def target_pad(s):
  while len(s) <= TARGET_LEN:
    s += ' '
  return s

  
# Generates one example (source, target) pair
def generate():
  a, b = number(), number()
  c = a * b
  source = source_pad('{0:b}*{0:b}'.format(a, b))
  target = target_pad('{0:b}'.format(c))
  assert all(ch in SOURCE_VOCAB for ch in source)
  assert all(ch in TARGET_VOCAB for ch in target)
  assert len(source) == SOURCE_LEN
  assert len(target) == TARGET_LEN
  return source, target

# Turns a string into a list of ints with the source embedding
def source_embed(s):
  return [SOURCE_VOCAB.index(ch) for ch in s]

# Turns a string into a list of ints with the target embedding
def target_embed(s):
  return [TARGET_VOCAB.index(ch) for ch in s]

  
class Model(object):

  '''
  Create a RNN sequence-to-sequence model for the problems created by 'generate'
  '''
  def __init__(self):

    # Set up hyperparameters
    self.num_layers = 2
    self.layer_size = 128

    # Set up the core RNN cells of the tensor network
    single_cell = rnn_cell.BasicLSTMCell(self.layer_size)
    self.cell = rnn_cell.MultiRNNCell([single_cell] * self.num_layers)

    # Set up placeholders for the inputs and outputs.
    # Leave batch size unspecified as a None shape.

    # The input problem
    self.encoder_inputs = [tf.placeholder(tf.int32,
                                          shape=[None],
                                          name='encoder{0}'.format(i))
                           for i in range(SOURCE_LEN)]

    # The correct answers
    self.labels = [tf.placeholder(tf.int32,
                                  shape=[None],
                                  name='labels{0}'.format(i))
                   for i in range(TARGET_LEN)]

    # Each item is equal, so weights are ones
    self.weights = [tf.ones_like(label, dtype=tf.float32)
                    for label in self.labels]

    # We will feed the decoder the correct output from the previous timestep,
    # with a "go" token on the first one
    self.decoder_inputs = [tf.placeholder(tf.int32,
                                          shape=[None],
                                          name='decoder{0}'.format(i))
                           for i in range(TARGET_LEN)]
    
    # Construct the guts of the model
    # For what exactly outputs and states are, see
    # https://github.com/tensorflow/tensorflow/blob/master/tensorflow/python/ops/seq2seq.py
    self.outputs, self.states = seq2seq.embedding_rnn_seq2seq(
      self.encoder_inputs,
      self.decoder_inputs,
      self.cell,
      len(SOURCE_VOCAB),
      len(TARGET_VOCAB))

    self.loss = seq2seq.sequence_loss(
      self.outputs,
      self.labels,
      self.weights)

    # Set up the ops we need for training
    learning_rate = 0.05
    momentum = 0.9
    self.optimizer = tf.train.MomentumOptimizer(learning_rate, momentum)
    self.train_op = optimizer.minimize(self.loss)

    self.sess = tf.Session()
    self.sess.run(tf.initialize_all_variables())

    
  '''
  Trains a batch of batch_size randomly generated problems.
  Returns the loss.
  '''
  def train_batch(self, batch_size):
    problems = [generate() for _ in range(batch_size)]

    encoder_input_data = np.transpose([source_embed(s) for s, _ in problems])
    label_data = np.transpose([target_embed(s) for _, s in problems])

    feed_dict = {}
    for i in range(SOURCE_LEN):
      feed_dict[self.encoder_inputs[i]] = encoder_input_data[i]
    for i in range(TARGET_LEN):
      feed_dict[self.labels[i]] = label_data[i]

    _, loss = self.sess.run([self.train_op, self.loss], feed_dict)
    return loss
    
