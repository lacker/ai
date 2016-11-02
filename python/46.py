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

# Prints if n is an odd composite that
# cannot be written as the sum of a prime and twice a square
def check_goldbach(n):
  if n % 2 == 0:
    return
  if is_prime(n):
    return
  for s in range(1, n):
    p = n - 2 * s * s
    if p <= 1:
      print n, "breaks goldbach"
      raise Exception('done')      
    if is_prime(p):
      print n, '=', p, '+ 2 *', s
      return

n = 3
while True:
  n += 2
  check_goldbach(n)
    
