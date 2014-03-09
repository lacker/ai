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
    i_x = 40
    i_y = 0
    j_x = 20
    j_y = 40
    for i in range(self.board.size):
      for j in range(self.board.size):
        x = i * i_x + j * j_x + 50
        y = i * i_y + j * j_y + 50
        color = ["black", "#EBCEAC", "white"][self.board.board[i][j] + 1]
        item_id = self.canvas.create_oval(x, y, x + 30, y + 30, fill=color)
        def onclick(event, a=i, b=j):
          self.click(a, b)
        self.canvas.tag_bind(item_id, "<ButtonPress-1>", onclick)
    self.root.update_idletasks()
        
  def click(self, i, j):
    # print "clicked", i, j
    for f in self.listeners:
      f((i, j))

if __name__ == "__main__":
  b = Board()
  v = Viewer(b)
  raise Exception("entering interpreter")
