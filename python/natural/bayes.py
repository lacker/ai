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
A distribution of things that don't have any relationship to each other.
They are just meaningless symbols.
In general None will never be a symbol - it is used to represent the
absence of a symbol.
'''
class SymbolDistribution(object):
  '''
  max_size is the maximum number of specific symbols for which we
  keep information around.
  history_range is the range of acceptable amounts of history.
  It's like a-b trees - when we hit the max we will rebalance to the min.
  '''
  def __init__(self, max_size=8, history_range=(1000, 1100)):
    self.history_range = history_range
    
    # Maps symbol to count
    self.symbol_count = {}
    
    # Things that are in total_count but not symbol_count are None
    self.total_count = 1

    
  def rebalance_if_needed(self):
    if self.total_count <= self.history_range[1]:
      return
      
    overshoot = self.history_range[1] / self.history_range[0]
    items = list(self.symbol_count.items())
    for s, c in items:
      self.symbol_count[s] = c / overshoot
    self.total_count /= overshoot

    
  '''
  Predicts the frequency of an unknown token, given an
  assumption about how frequent it is.
  (If we have no assumption, the prediction is simply zero.)
  '''
  def predict_unknown(self, assumption):
    min_count = min(self.symbol_count.values())

    estimation = (min_count - 1 / overshoot)
    return min(estimation, assumption)

  '''
  Predicts the frequency of a given symbol.
  '''
  def predict(self, symbol):
    return self.symbol_count.get(symbol, 0) / self.total_count
  
  '''
  Merges this symbol distribution with another one.
  Optimizes the resulting symbol distribution, given the assumption
  that these reflect the same underlying distribution.
  '''
  def merge(self, other):
    raise 'TODO'
