#!/usr/bin/python

from math import sqrt

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

answer = 1

for p in range(1000):
  if is_prime(p):
    print answer
    answer *= p
    if answer > 1000000:
      break


