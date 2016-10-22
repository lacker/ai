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

def iter_families(num_digits):
  if num_digits < 1:
    yield ''
    return

  for char in 'x0123456789':
    for rest in iter_families(num_digits - 1):
      yield char + rest

def search_family(family):
  if 'x' not in family:
    return 0
  primes = []
  for char in '0123456789':
    s = family.replace('x', char)
    if s[0] != '0' and is_prime(int(s)):
      primes.append(s)
  if len(primes) >= 7:
    print len(primes), ', '.join(primes)
      
def search_families(num_digits):
  print 'search_families', num_digits
  for fam in iter_families(num_digits):
    search_family(fam)
  
  
search_families(5)
search_families(6)
search_families(7)
