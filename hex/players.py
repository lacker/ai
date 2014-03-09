#!/usr/bin/python -i
"""
Some hex-playing strategies.
"""

from collections import defaultdict
import random
import time

from hex import board
from hex import viewer


"Picks a move at random."
def rand(b):
  return random.choice(b.empty_spots())

def partcrazy(b):
  if random.random() < 0.1:
    return rand(b)
  return montecarlo(b)
  
"""
Does treeless RAVE algorithm.
"""
def montecarlo(b):
  mover_wins = defaultdict(int)
  mover_losses = defaultdict(int)
  empty = b.empty_spots()
  mover_score = 0
  mover = b.to_move
  
  playouts = 2000
  for n in range(playouts):
    random.shuffle(empty)
    copy = b.copy()
    for spot in empty:
      copy.move(spot, check_for_winner=False)
    winner = copy.winner()
    if winner == mover:
      mover_score += 1
      for i, j in empty:
        if copy.board[i][j] == mover:
          mover_wins[(i, j)] += 1
    else:
      for i, j in empty:
        if copy.board[i][j] == mover:
          mover_losses[(i, j)] += 1

  ranked = []
  for spot in empty:
    wins = mover_wins[spot]
    total = wins + mover_losses[spot]
    predicted = (1.0 + wins) / (1.0 + total)
    ranked.append((predicted, spot))
  score, spot = max(ranked)
  print "current score: %d/%d = %.2f" % (mover_score, playouts,
                                         mover_score / float(playouts))
  print "%s score: %d/%d ~= %.2f" % (
    spot,
    mover_wins[spot],
    mover_wins[spot] + mover_losses[spot],
    score)
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

  make_computer(montecarlo, board.BLACK, b)

  make_computer(partcrazy, board.WHITE, b)
  # make_human(v, board.WHITE, b)

  v.root.after_idle(lambda: b.move((0, 1)))
