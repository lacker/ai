#!/usr/bin/python -i
"""
A UI for watching a hex game or playing one.
"""

from Tkinter import *

from hex.board import Board

class Viewer(object):
  def __init__(self, board):
    self.board = board

    self.root = Tk()
    self.canvas = Canvas(self.root, width=750, height=550)
    self.canvas.pack()

    self.listeners = []
    
    self.redraw()
    self.board.add_listener(lambda: self.redraw())

  def add_listener(self, listener):
    self.listeners.append(listener)

  def spot_center(self, row, col):
    col_x = 40
    col_y = 0
    row_x = 20
    row_y = 40

    x = row * row_x + col * col_x + 65
    y = row * row_y + col * col_y + 65
    return x, y
    
  def redraw(self):
    # print "redrawing"
    self.canvas.delete(ALL)
    self.canvas.create_rectangle(3, 3, 750, 550, fill="#EBCEAC")

    a, b = self.spot_center(0, 5)
    c, d = self.spot_center(10, 5)
    self.canvas.create_line(a, b, c, d, fill="black")
    a, b = self.spot_center(5, 0)
    c, d = self.spot_center(5, 10)
    self.canvas.create_line(a, b, c, d, fill="black")
    
    for row in range(self.board.size):
      for col in range(self.board.size):
        x, y = self.spot_center(row, col)
        color = ["black", "#EBCEAC", "white"][self.board.board[row][col] + 1]
        item_id = self.canvas.create_oval(x - 15, y - 15,
                                          x + 15, y + 15, fill=color)
        def onclick(event, r=row, c=col):
          self.click(r, c)
        self.canvas.tag_bind(item_id, "<ButtonPress-1>", onclick)

    self.root.update_idletasks()
        
  def click(self, r, c):
    # print "clicked", r, c
    for f in self.listeners:
      f((r, c))

if __name__ == "__main__":
  b = Board()
  v = Viewer(b)
  raise Exception("entering interpreter")
