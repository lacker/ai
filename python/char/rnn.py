#!/usr/bin/python
"""
Character-level recurrent neural networks, using Keras.
See:
https://github.com/karpathy/char-rnn
for inspiration.
"""

import os

def make_labeled_data(chars):
  """
  Given a list of characters, turn it into labeled data.
  That involves turning all characters into numbers.
  In general the input vector is all characters leading up to a
  certain point, and the target label is the next character.
  """
  dataset, labels = [], []
  ints = map(ord, chars)
  accum = []
  for i in ints:
    dataset.append(list(accum))
    labels.append(i)
    accum.append(i)
  return dataset, labels
  

chatfile = os.path.realpath(__file__ + "/../data/chat.txt")
print "chatfile:", chatfile
