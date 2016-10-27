#!/usr/bin/env python

from itertools import permutations

def is_prime(n):
  if n < 2:
    return False
  if n == 2:
    return True
  if n % 2 == 0:
    return False
  d = 3
  while True:
    if d * d > n:
      return True
    if n % d == 0:
      return False
    d += 2

def pandigitals(n):
  for perm in permutations(range(1, n + 1)):
    yield int(''.join(map(str, perm)))

for n in range(1, 10):
  for pan in pandigitals(n):
    if is_prime(pan):
      print pan
