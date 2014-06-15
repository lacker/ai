#!/usr/bin/python -i
"""
A hex board class useful for basic operations.
"""

import json

BLACK = -1
WHITE = 1
EMPTY = 0

def color_name(color):
  return ["Black", "Empty", "White"][color + 1]

class Board(object):
  def __init__(self, size=11):
    self.size = size

    # The board is a grid. Each value is either BLACK, WHITE, or EMPTY.
    # Cells are typically referred to with a letter plus a number.
    # Display would look like a diamond, e.g. for a 4x4 board:
    #
    #  ABCD
    # -----        
    # \    \       1
    #  \    \      2
    #   \    \     3
    #    \    \    4
    #     -----
    #
    # The board is indexed by (row, column). So the first index
    # represents the integer, and the second index represents the
    # letter.
    # Black goes top to bottom; White goes left to right.
    
    self.board = [[EMPTY] * self.size for _ in range(self.size)]

    # Let's say black goes first
    self.to_move = BLACK

    # Track move history
    self.history = []
    
    self.listeners = []


  """
  Converts to json.
  Uses the keys "Board" and "ToMove" for go compatibility.
  """
  def to_json(self):
    return json.dumps({
      "Board": self.board,
      "ToMove": self.to_move,
      })
    
  """
  Copies the board. If transpose, transposes to replace black and
  white.
  Does not copy the move history.
  """
  def copy(self, transpose=False):
    new_board = Board(self.size)
    if transpose:
      new_board.to_move = -self.to_move
    else:
      new_board.to_move = self.to_move
    for i in range(self.size):
      for j in range(self.size):
        if transpose:
          new_board.board[i][j] = -self.board[j][i]
        else:
          new_board.board[i][j] = self.board[i][j]
    return new_board

  "Shortcut for copying with transpose=True."
  def transpose(self):
    return self.copy(transpose=True)
    
  "Returns the empty spots."
  def empty_spots(self):
    return [(i, j) for i in range(self.size) for j in range(self.size)
            if self.board[i][j] == EMPTY]

  "Adds a listener to get triggered whenever a move is made."
  def add_listener(self, listener):
    self.listeners.append(listener)
    
  """
  Makes a move. Returns whether that move was valid.
  If the move was valid, triggers listeners.
  """
  def move(self, spot, check_for_winner=True):
    i, j = spot
    if i < 0 or i >= self.size or j < 0 or j >= self.size:
      return False
    prev = self.board[i][j]
    if prev != EMPTY:
      return False
    self.board[i][j] = self.to_move
    self.to_move = -self.to_move

    if check_for_winner:
      winner = self.winner()
      if winner != EMPTY:
        self.to_move = EMPTY
        print color_name(winner), "wins!"

    self.history.append(spot)
        
    for f in self.listeners:
      f()
        
    return True

  "Return whether black has won the game."
  def did_black_win(self):
    # Searches top-down. Start from [0][_] which is the top row.
    active = set()
    checked = set()
    for i in range(self.size):
      if self.board[0][i] == BLACK:
        active.add((0, i))

    while active:
      spot = active.pop()
      checked.add(spot)

      # Find all the neighboring black stones
      i, j = spot
      for a, b in [(i - 1, j),
                   (i + 1, j),
                   (i, j + 1),
                   (i, j - 1),
                   (i + 1, j - 1),
                   (i - 1, j + 1)]:
        if a < 0 or a >= self.size or b < 0 or b >= self.size:
          continue
        if self.board[a][b] != BLACK:
          continue
        new_spot = (a, b)
        if new_spot in checked:
          continue
        if a == self.size - 1:
          return True
        active.add(new_spot)

    return False

  "Return who won, or EMPTY if neither."
  def winner(self):
    if self.did_black_win():
      return BLACK
    if self.transpose().did_black_win():
      return WHITE
    return EMPTY
