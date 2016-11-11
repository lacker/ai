#!/usr/bin/env python

n = 0
counter = {}
smallest = {}
digits = 0
maxcount = 0
while True:
  n += 1
  cube = n ** 3
  key = ''.join(sorted(list(str(cube))))
  if len(key) > digits:
    digits = len(key)
    if maxcount >= 5:
      break
  if key not in counter:
    smallest[key] = cube
    counter[key] = 0
    
  counter[key] += 1
  if counter[key] >= 4:
    maxcount = max(maxcount, counter[key])
    print n, 'cubed is', cube, 'with count', counter[key], 'for smallest', smallest[key]
