#!/usr/bin/env python

from itertools import combinations

'''
Create a list of all cubes.
A cube is a sorted size-6 subset of 0 1 2 3 4 5 6 7 8 9
'''
CUBES = map(tuple, combinations(range(10), 6))

def has_digit(c, d):
  if d in c:
    return True
  if d == 9 and 6 in c:
    return True
  if d == 6 and 9 in c:
    return True
  return False

  
'''
Checks if two cubes cover the digits provided
'''
def covers(c1, c2, d1, d2):
  if has_digit(c1, d1) and has_digit(c2, d2):
    return True
  if has_digit(c1, d2) and has_digit(c2, d1):
    return True
  return False

def coverall(c1, c2):
  return (covers(c1, c2, 0, 1) and
          covers(c1, c2, 0, 4) and
          covers(c1, c2, 0, 9) and
          covers(c1, c2, 1, 6) and
          covers(c1, c2, 2, 5) and
          covers(c1, c2, 3, 6) and
          covers(c1, c2, 4, 9) and
          covers(c1, c2, 6, 4) and
          covers(c1, c2, 8, 1))

count = 0
for c1, c2 in combinations(CUBES, 2):
  if coverall(c1, c2):
    print c1, c2
    count += 1
print count
