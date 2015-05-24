#!/usr/bin/python
"""
Character-level recurrent neural networks, using Keras.
See:
https://github.com/karpathy/char-rnn
for inspiration.
"""

import math
import os
import random

def random_number():
  "Makes a random number with 5 to 20 digits."
  return int(10 ** (5 + (random.random() * 15)))

def mod_statement(m):
  """
  Makes a random expression that calculates mod m.
  For example, a mod_statement(3) could be:
  "8374 % 3 = 1"
  """
  num = random_number()
  rhs = num % m
  return "%d %% %d = %d" % (num, m, rhs)
  
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
