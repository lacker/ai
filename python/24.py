#!/usr/bin/env python

from itertools import permutations

count = 1
for perm in permutations('0123456789'):
  if count == 1000000:
    print ''.join(perm)
    break
  count += 1
