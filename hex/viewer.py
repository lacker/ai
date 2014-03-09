#!/usr/bin/python -i
"""
A UI for watching a hex game or playing one.
"""

class Viewer(object):
  def __init__(self, board):
    self.board = board

    self.redraw()
    self.board.add_listener(lambda: self.redraw())

  def redraw(self):
    pass
