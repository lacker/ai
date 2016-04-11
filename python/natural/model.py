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

SOURCE_VOCAB, TARGET_VOCAB = '10*\n', '10\n'
MAX_SOURCE_LEN = 2 + 2 * len(str(MAX_NUMBER))
MAX_TARGET_LEN = 5

# Generates one example (source, target) pair
def generate():
    a, b = number(), number()
    c = a * b
    source, target = '{0:b}*{0:b}\n'.format(a, b), '{0:b}\n'.format(c)
    assert all(ch in SOURCE_VOCAB for ch in source)
    assert all(ch in TARGET_VOCAB for ch in target)
    assert len(source) <= MAX_SOURCE_LEN
    assert len(target) <= MAX_TARGET_LEN
    return source, target

    
class Model(object):

  '''
  If learning=False, that means we are not training, and we don't need to
  learn things on the fly.
  '''
  def __init__(self, learning=True):

    # Set up hyperparameters
    self.num_layers = 2
    self.layer_size = 128

    # Set up the core RNN cells of the tensor network
    self.single_cell = rnn_cell.BasicLSTMCell(self.layer_size)
    self.multi_cell = rnn_cell.MultiRNNCell(
      [self.single_cell] * self.num_layers)

    # Set up placeholders for the source and target embeddings
    self.encoder_inputs = [tf.placeholder(tf.int32,
                                          shape=[None],
                                          name='encoder{0}'.format(i))
                           for i in range(len(SOURCE_VOCAB))]
    
    self.decoder_inputs = [tf.placeholder(tf.int32,
                                          shape=[None],
                                          name='decoder{0}'.format(i))
                           for i in range(len(TARGET_VOCAB))]
    
    self.target_weights = [tf.placeholder(tf.float32,
                                          shape=[None],
                                          name='weight{0}'.format(i))
                           for i in range(len(TARGET_VOCAB))]    
    
    # Our targets are decoder inputs shifted by one.
    self.targets = self.decoder_inputs[1:]


