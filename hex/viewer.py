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
    self.canvas.create_rectangle(3, 3, 750, 550, fill="#EBCEAC")

    self.listeners = []
    
    self.redraw()
    self.board.add_listener(lambda: self.redraw())

  def add_listener(self, listener):
    self.listeners.append(listener)
  
  def redraw(self):
    print "redrawing"
    col_x = 40
    col_y = 0
    row_x = 20
    row_y = 40
    for row in range(self.board.size):
      for col in range(self.board.size):
        x = row * row_x + col * col_x + 50
        y = row * row_y + col * col_y + 50
        color = ["black", "#EBCEAC", "white"][self.board.board[row][col] + 1]
        item_id = self.canvas.create_oval(x, y, x + 30, y + 30, fill=color)
        def onclick(event, r=row, c=col):
          self.click(r, c)
        self.canvas.tag_bind(item_id, "<ButtonPress-1>", onclick)
    self.root.update_idletasks()
        
  def click(self, r, c):
    print "clicked", r, c
    for f in self.listeners:
      f((r, c))

if __name__ == "__main__":
  b = Board()
  v = Viewer(b)
  raise Exception("entering interpreter")
