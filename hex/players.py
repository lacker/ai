#!/usr/bin/python -i
"""
Some hex-playing strategies.
"""

import core.hex.board.Board as Board
import random

"Picks a move at random."
def rand(board):
  return random.choice(board.empty_spots())

"Displays a board for a human to choose a move."
def human(board):
  raise Exception("not implemented yet")
