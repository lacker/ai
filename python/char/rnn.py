#!/usr/bin/python
"""
Character-level recurrent neural networks, using Keras.
See:
https://github.com/karpathy/char-rnn
for inspiration.

Philosophically training on deterministic problems with which we have
infinite training data seems interesting.

State of the art is roughly, it works for addition, but only for a
fixed length of sequence, and it doesn't work for anything more
complex. I wonder if a new type of neuron might be able to do
something like a stack data structure.

TODO: make the GPU work
"""

import math
import os
import random

from keras.models import Sequential

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

def make_labeled_mod3(n):
  """
  Makes a labeled dataset of n mod3 problems.
  We can take it for granted that a \n resets and not ask the machine
  to learn that. It only has to learn that \n terminates the correct
  answer, and we'll take it from there.
  """
  dataset, labels = [], []
  for _ in range(n):
    statement = mod_statement(3) + "\n"
    newdata, newlabels = make_labeled_data(statement)
    dataset.extend(newdata)
    labels.extend(newlabels)
  return dataset, labels


# Let's use this dataset
mod3_data, mod3_labels = make_labeled_mod3(10000)

print "creating a model"
model = Sequential()
