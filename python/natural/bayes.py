#!/usr/bin/env python3


'''
One node in a graph that does a Bayesian modeling of a problem.
'''
class VoidToBool(object):
  '''
  max_history is the point at which we stabilize our learning rate
  '''
  def __init__(self, max_history=1000):
    self.yes_count = 0.5
    self.no_count = 0.5
    self.max_history = max_history

  '''
  Updates the internal model given one instance where the correct output is 'target'.
  '''
  def feedback(self, target):
    if target:
      self.yes_count += 1
    else:
      self.no_count += 1

    overshoot = (self.yes_count + self.no_count) / self.max_history
    if overshoot > 1:
      self.yes_count /= overshoot
      self.no_count /= overshoot
      
  '''Predicts the odds of "yes".'''
  def predict_yes(self):
    return self.yes_count / (self.yes_count + self.no_count)


'''
A generic predictor that works off no input.
'''
class VoidInput(object):
  def __init__(self, default='default', max_history=1000):
    raise 'TODO'
