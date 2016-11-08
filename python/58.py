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

# Initialize with stats from sidelen = 1
primes = 0
total = 1
n = 1

while True:
  n += 2
  diff = (n - 1)
  se = n * n
  for k in (se, se - diff, se - 2 * diff, se - 3 * diff):
    total += 1
    if is_prime(k):
      print k, 'is prime'
      primes += 1

  print 'SO FAR', n, primes, total
  if float(primes) / total < 0.1:
    raise Exception('done')
