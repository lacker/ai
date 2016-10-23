#!/usr/bin/env python

def make_number_spiral(n):
  if n == 1:
    return [[1]]

  center = make_number_spiral(n - 2)

  upper_right = n * n
  upper_left = upper_right - (n - 1)
  bottom_left = upper_left - (n - 1)
  bottom_right = bottom_left - (n - 1)
  virtual_upper_right = bottom_right - (n - 1)

  answer = []
  for y in range(n):
    line = []
    answer.append(line)
    for x in range(n):
      if y == 0:
        line.append(upper_left + x)
      elif y == n - 1:
        line.append(bottom_left - x)
      elif x == 0:
        line.append(upper_left - y)
      elif x == n - 1:
        line.append(virtual_upper_right + y)
      else:
        line.append(center[y - 1][x - 1])
  return answer

def mprint(spiral):
  for line in spiral:
    print line

def sum_diag(spiral):
  answer = 0
  n = len(spiral)
  for i in range(n):
    for j in set([i, n - i - 1]):
      answer += spiral[i][j]
  return answer

def sum_diag_spiral(n):
  if n == 1:
    return 1
  upper_right = n * n
  upper_left = upper_right - (n - 1)
  bottom_left = upper_left - (n - 1)
  bottom_right = bottom_left - (n - 1)
  return sum_diag_spiral(n - 2) + upper_right + upper_left + bottom_left + bottom_right


  
print sum_diag_spiral(5)
print sum_diag_spiral(1001)
    
