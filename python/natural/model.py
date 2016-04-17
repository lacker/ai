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
# 001+001
# 0001000
# least-significant-bit first.
# zero-padded on the right to get consistent lengths.
SOURCE_VOCAB, TARGET_VOCAB = '10+', '10'
NUMBER_LEN = len('{0:b}'.format(MAX_NUMBER))
SOURCE_LEN = 1 + 2 * NUMBER_LEN
TARGET_LEN = len('{0:b}'.format(MAX_NUMBER + MAX_NUMBER))

# Converts to LSB int if needed
def right_pad(s, ch, length):
  if type(s) == int:
    s = '{0:b}'.format(s)[::-1]
  while len(s) < length:
    s += ch
  return s

  
# Generates one example (source, target) pair
def generate():
  a, b = number(), number()
  c = a + b
  a_str = right_pad(a, '0', NUMBER_LEN)
  b_str = right_pad(b, '0', NUMBER_LEN)
  source = a_str + '+' + b_str
  target = right_pad(c, '0', TARGET_LEN)
  assert all(ch in SOURCE_VOCAB for ch in source)
  assert all(ch in TARGET_VOCAB for ch in target)
  assert len(source) == SOURCE_LEN, 'source: [{}] #{}'.format(source, SOURCE_LEN)
  assert len(target) == TARGET_LEN, 'target: [{}] #{}'.format(target, TARGET_LEN)
  return source, target

  
# Turns a string into a list of ints with the source embedding
def source_embed(s):
  return [SOURCE_VOCAB.index(ch) for ch in s]

# Turns a string into a list of ints with the target embedding
def target_embed(s):
  return [TARGET_VOCAB.index(ch) for ch in s]

# Reverses target_embed
def target_unembed(data):
  return ''.join(TARGET_VOCAB[i] for i in data)

  
class Model(object):

  '''
  Create a RNN sequence-to-sequence model for the problems created by 'generate'
  '''
  def __init__(self):

    # Set up hyperparameters
    self.num_layers = 3
    self.layer_size = 256

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

    # decoder_inputs has the correct output from the previous timestep,
    # with a zero-hot "go" token on the first one
    go_token = tf.zeros_like(self.labels[0], dtype=np.int32, name="GO")
    self.decoder_inputs = [go_token] + self.labels[:-1]
    
    # Construct the guts of the model.
    # This same model will be used for training and testing, so we
    # don't feed_previous.
    self.outputs, self.states = seq2seq.embedding_rnn_seq2seq(
      self.encoder_inputs,
      self.decoder_inputs,
      self.cell,
      len(SOURCE_VOCAB),
      len(TARGET_VOCAB),
      feed_previous=False)

    self.loss = seq2seq.sequence_loss(
      self.outputs,
      self.labels,
      self.weights)

    # Set up the ops we need for training

    if True: # momentum
      learning_rate = 0.05
      momentum = 0.9
      self.optimizer = tf.train.MomentumOptimizer(learning_rate, momentum)
      self.train_op = self.optimizer.minimize(self.loss)
    else: # adam
      # Assumes batch size of 100
      self.cost = tf.reduce_sum(self.loss) / TARGET_LEN / 100
      self.lr = tf.Variable(0.0, trainable=False)
      tvars = tf.trainable_variables()
      # Clip gradients at 5.0
      grads, _ = tf.clip_by_global_norm(tf.gradients(self.cost, tvars),
                                        5.0)
      optimizer = tf.train.AdamOptimizer(self.lr)
      self.train_op = optimizer.apply_gradients(zip(grads, tvars))

    self.sess = tf.Session()
    self.sess.run(tf.initialize_all_variables())

    
  '''
  Trains a batch of batch_size randomly generated problems.
  Returns the loss.
  '''
  def train_batch(self, batch_size):
    problems = [generate() for _ in range(batch_size)]

    input_data = np.transpose([source_embed(s) for s, _ in problems])
    label_data = np.transpose([target_embed(s) for _, s in problems])

    feed_dict = {}
    for i in range(SOURCE_LEN):
      feed_dict[self.encoder_inputs[i]] = input_data[i]
    for i in range(TARGET_LEN):
      feed_dict[self.labels[i]] = label_data[i]

    _, loss = self.sess.run([self.train_op, self.loss], feed_dict)
    return loss

    
  '''
  Tests a batch of batch_size randomly generated problems.
  Returns a list of (input, label, output) tuples.
  '''
  def test_batch(self, batch_size):
    problems = [generate() for _ in range(batch_size)]
    inputs = [i for i, _ in problems]
    labels = [l for _, l in problems]
    
    input_data = np.transpose([source_embed(s) for s, _ in problems])
    label_data = np.transpose([target_embed(s) for _, s in problems])

    feed_dict = {}
    for i in range(SOURCE_LEN):
      feed_dict[self.encoder_inputs[i]] = input_data[i]
    for i in range(TARGET_LEN):
      feed_dict[self.labels[i]] = label_data[i]
      
    output_data = self.sess.run(self.outputs, feed_dict)
    unembedded = np.transpose(np.argmax(np.array(output_data), axis=2))
    outputs = [target_unembed(data) for data in unembedded]

    return zip(inputs, labels, outputs)


  def score(self, n=1000):
    success = 0
    total = 0
    for inp, label, out in self.test_batch(n):
      total += 1
      if label == out:
        success += 1
    print('{} / {} = {}%'.format(success, total, 100.0 * success / total))


  def sample(self, n=10):
    for inp, label, out in self.test_batch(n):
      print('input puzzle: {}'.format(inp))
      print('right answer: {}'.format(label))
      if label == out:
        print('OK')
      else:
        print('wrong answer: {}'.format(out))
      print()
