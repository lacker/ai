#!/usr/bin/python -i
"""
A hex board class useful for basic operations.
"""

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
    # In the board variable, the first index represents the number,
    # and the second the integer.
    # Black goes top to bottom; White goes left to right.
    
    self.board = [[EMPTY] * self.size for _ in range(self.size)]

    # Let's say black goes first
    self.to_move = BLACK

    self.listeners = []

    
  """
  Constructs a new board that exchanges the role of black and white.
  """
  def transpose(self):
    new_board = Board(self.size)
    new_board.to_move = -self.to_move
    for i in range(self.size):
      for j in range(self.size):
        new_board.board[i][j] = -self.board[j][i]
    return new_board

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
  def move(self, spot):
    i, j = spot
    if i < 0 or i >= self.size or j < 0 or j >= self.size:
      return False
    prev = self.board[i][j]
    if prev != EMPTY:
      return False
    self.board[i][j] = self.to_move
    self.to_move = -self.to_move
    for f in self.listeners:
      f()

    winner = self.winner()
    if winner != EMPTY:
      self.to_move = EMPTY
      print color_name(winner), "wins!"
    return True

  "Return whether black has won the game."
  def did_black_win(self):
    # Searches top-down. Start from [_][0]
    active = set()
    checked = set()
    for i in range(self.size):
      if self.board[i][0] == BLACK:
        active.add((i, 0))

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
        if b == self.size - 1:
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
