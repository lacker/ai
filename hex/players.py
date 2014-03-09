#!/usr/bin/python -i
"""
Some hex-playing strategies.
"""

from hex import board
from hex import viewer
import random

"Picks a move at random."
def rand(b):
  return random.choice(b.empty_spots())

"""
Treats a human's clicks as playing for this color whenever it is the
appropriate turn.
"""
def make_human(viewer, color, b):
  def try_to_play(spot):
    if b.to_move != color:
      return
    b.move(spot)
      
  viewer.add_listener(try_to_play)
  
"""
Makes a computer player with the given function that will decide what
to play.
"""
def make_computer(f, color, b):
  def play_if_my_turn():
    if b.to_move == color:
      b.move(f(b))
  b.add_listener(play_if_my_turn)
  
if __name__ == "__main__":
  b = board.Board()
  v = viewer.Viewer(b)

  # Make the human play vs rand
  make_computer(rand, board.BLACK, b)
  make_human(v, board.WHITE, b)

  b.move((1, 1))
