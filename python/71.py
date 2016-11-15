#!/usr/bin/python

from math import floor

best_ratio = 0.0

for d in range(1, 1000001):
  if d % 7 == 0:
    continue
  n = int(floor((3./7) * d))
  ratio = n / float(d)
  if ratio > best_ratio:
    print n, '/', d, '=', ratio
    best_ratio = ratio
