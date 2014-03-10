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
Does a tree on top of monte carlo where the tree is just a single
step - one move, then the responder move.
"""
def shallow_tree(b):
  empty = b.empty_spots()
  mover = b.to_move
  # Record maps (our spot, their spot) to the win/loss record.
  record = {}

  # Run the playouts
  playouts = 2000
  for n in range(playouts):
    random.shuffle(empty)
    copy = b.copy()
    for spot in empty:
      copy.move(spot, check_for_winner=False)
    winner = copy.winner()
    if winner == mover:
      dwin, dloss = 1, 0
    else:
      dwin, dloss = 0, 1
    our_spots = []
    their_spots = []
    for i, j in empty:
      if copy.board[i][j] == mover:
        our_spots.append((i, j))
      else:
        their_spots.append((i, j))
    for our_spot in our_spots:
      for their_spot in their_spots:
        key = (our_spot, their_spot)
        w, l = record.get(key, (0, 0))
        record[key] = (w + dwin, l + dloss)
        
  # Score each of our moves against the worst-performing response
  ranked = []
  for our_spot in empty:
    worst_score = 1.0
    for their_spot in empty:
      if our_spot == their_spot:
        continue
      w, l = record.get((our_spot, their_spot), (0, 0))
      possible_score = (1.0 + w) / (1.0 + w + l)
      worst_score = min(worst_score, possible_score)
    ranked.append((worst_score, our_spot))

  # Pick our best move
  score, spot = max(ranked)
  print "best move: %s. score: %.2f" % (spot, score)
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

  make_computer(shallow_tree, board.BLACK, b)

  # make_computer(montecarlo, board.WHITE, b)
  make_human(v, board.WHITE, b)

  v.root.after_idle(lambda: b.move((3, 3)))
