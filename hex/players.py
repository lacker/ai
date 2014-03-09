#!/usr/bin/python -i
"""
Some hex-playing strategies.
"""

from collections import defaultdict
import random

from hex import board
from hex import viewer


"Picks a move at random."
def rand(b):
  return random.choice(b.empty_spots())

"""
Does treeless RAVE algorithm.
I.e., in random playouts, picks the cell that is most likely to
correspond to the winning side.
"""
def montecarlo(b):
  scores = defaultdict(int)
  empty = b.empty_spots()

  for n in range(50):
    random.shuffle(empty)
    copy = b.copy()
    for spot in empty:
      copy.move(spot, check_for_winner=False)
    winner = copy.winner()
    for i, j in empty:
      if copy.board[i][j] == winner:
        scores[(i, j)] += 1

  print scores
        
  score, spot = max((sc, sp) for (sp, sc) in scores.iteritems())
  print "best score: %d at %s" % (score, spot)
  return spot

  
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

  # Make the human play vs a computer
  make_computer(montecarlo, board.BLACK, b)
  make_human(v, board.WHITE, b)

  b.move((1, 1))
