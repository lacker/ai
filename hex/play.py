#!/usr/bin/python -i
"""
Some hex-playing strategies.
"""

from collections import defaultdict
import json
import os
import random
import subprocess
import sys
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
Constructs a go player that will shell to go with the given player type
"""
def go_player(player_type):
  return lambda b: go_shell(player_type, b)
  
"""
Shells out to go to play a move.
"""
def go_shell(player_type, b):
  fname = board.__file__ + "/../../go/src/lacker.info/play_hex.go"
  fname = os.path.abspath(fname)
  output = subprocess.check_output([
    "go", "run", fname, player_type, b.to_json()])
  json_spot = json.loads(output)
  answer = json_spot["Row"], json_spot["Col"]
  print player_type, "played", answer
  return answer
  
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
  playouts = 500
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
    print "human played", spot
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


"""
Runs a sample game starting with a random move.
Must run two computers.
Returns the board.
"""
def sample_game(white_player, black_player):
  b = board.Board()
  make_computer(black_player, board.BLACK, b)
  make_computer(white_player, board.WHITE, b)
  b.move((random.randint(0, 10), random.randint(0, 10)))
  return b

"""
Saves a game history to a csv file. Appends it.
"""
def save_history(b, filename="games.csv"):
  encoded = ",".join(["%d-%d" % spot for spot in b.history])
  f = open(os.path.expanduser("~/data/" + filename), "a")
  f.write(encoded + "\n")
  f.close()
  
def make_player_by_type(viewer, color, board, player_type):
  if player_type == "human":
    return make_human(viewer, color, board)
  return make_computer(go_player(player_type), color, board)

  
if __name__ == "__main__":
  # Usage:
  # ./play.py blackplayer whiteplayer numgames
  # For example:
  # ./play.py sr5 human 3
  
  # Run a game that we watch.
  # Must be run with -i
  b = board.Board()
  v = viewer.Viewer(b)

  players = list(sys.argv[1:3])
  while len(players) < 2:
    players.append("human")

  num_games = 1
  if len(sys.argv) >= 4:
    num_games = int(sys.argv[3])

  wins = {}
  games_played = 0
  
  print "playing", players[0], "vs", players[1]

  first_move = (1, 1)
  
  make_player_by_type(v, board.BLACK, b, players[0])
  make_player_by_type(v, board.WHITE, b, players[1])

  # Starts a new game after this one is over
  def check_for_win():
    if b.to_move == board.EMPTY:
      winner = b.winner()
      wins[winner] = wins.get(winner, 0) + 1
      print "%s (Black): %d - %s (White): %d" % (
        players[0], wins.get(board.BLACK, 0),
        players[1], wins.get(board.WHITE, 0))
      if num_games > sum(wins.values()):
        b.reset()
        v.root.after_idle(lambda: b.move(first_move))
      
  b.add_listener(check_for_win)
  
  v.root.after_idle(lambda: b.move(first_move))
