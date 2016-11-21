#!/usr/bin/env python

from math import sqrt

'''
Yields all a such that (a, b, c) is a pythagorean triple with a < b
and a <= m
'''
def triangles(b, m):
  for a in range(3, min(b, m + 1)):
    c = sqrt(a*a + b*b)
    if int(c) == c:
      yield a

'''
Counts the number of cuboids with int shortest-corner-path with this b.

IE cases where a^2 + b^2 is a square defining the cuboid-shortest-path
and a < b.
'''
def bcount(b, m):
  answer = 0
  
  for a in triangles(b, m):
    # If b <= m, the two shortest sides could add to a
    if b <= m:
      answer += a / 2

    # We need to count (c, d) pairs with c, d <= a and c + d == b
    # c ranges in [b-a, a]
    if b - a > a:
      continue
    answer += (2 * a - b + 2) / 2

  return answer

def cuboidcount(m):
  answer = 0
  for b in range(4, 2 * m):
    answer += bcount(b, m)
  return answer

for m in range(1815, 100000):
  cc = cuboidcount(m)
  print m, cc
  if cc > 1000000:
    break
