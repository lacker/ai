#!/usr/bin/env python

from itertools import permutations



def find():
  answers = set()
  for perm in permutations('123456789'):
    for lenx in range(1, 8):
      for leny in range(1, 9 - lenx):
        x = int(''.join(perm[:lenx]))
        y = int(''.join(perm[lenx:lenx + leny]))
        if y > x:
          break
        z = int(''.join(perm[lenx+leny:]))
        if x * y > z:
          break
        if x * y == z:
          print x, y, z
          if z not in answers:
            answers.add(z)
  print sum(answers)
  
find()
