#!/usr/bin/env python

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

primes = set()
for k in range(1000000):
  if is_prime(k):
    print 'adding', k
    primes.add(str(k))
    
def rotate(s):
  return s[1:] + s[0]

def check_all_rot(s):
  original = s
  while s in primes:
    s = rotate(s)
    if s == original:
      return True
  return False

answer = 0
for k in range(1000000):
  if check_all_rot(str(k)):
    print k
    answer += 1
print answer
