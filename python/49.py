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

anagrams = {}

def anagram_key(n):
  return ''.join(sorted(list(str(n))))

primes = set()
for k in range(10000):
  if is_prime(k):
    print 'adding prime:', k
    primes.add(k)
    key = anagram_key(k)
    if key not in anagrams:
      anagrams[key] = []
    anagrams[key].append(k)


for key, value in anagrams.items():
  if len(value) < 3:
    continue
  if len(key) != 4:
    continue
  for a in value:
    for b in value:
      if b <= a:
        continue
      c = 2 * b - a
      if c in value:
        print a, b, c
