#!/usr/bin/env python3


'''
One node in a graph that does a Bayesian modeling of a problem.
'''
class VoidToBool(object):
  def __init__(self):
    self.yes_count = 0
    self.no_count = 0

  '''Reports one "yes" of training data.'''
  def train_yes(self):
    self.yes_count += 1

  '''Reports one "no" of training data.'''
  def train_no(self):
    self.no_count += 1

  '''Predicts the odds of "yes".'''
  def predict(self):
    return (0.5 + self.yes_count) / (0.5 + self.no_count)
