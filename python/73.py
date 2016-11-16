#!/usr/bin/env python

from math import floor, ceil

def gcd(a, b):
  if a < 0:
    a = -a
  if b < 0:
    b = -b
  if a < b:
    a, b = b, a
  if b == 0:
    return a
  return gcd(b, a % b)

def runupto(x):
  answer = 0
  for d in range(4, x + 1):
    # Find where n/d = 1/3 rounded up
    low_n = int(ceil(d / 3.0 + 0.0000001))

    # Find where n/d = 1/3 rounded down
    high_n = int(floor(d / 2.0 + 0.000000001))

    for n in range(low_n, high_n + 1):
      if gcd(n, d) == 1:
        print n, '/', d
        answer += 1
  return answer

print runupto(12000)
