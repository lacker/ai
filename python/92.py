#!/usr/bin/env python

results = {}

def solve(n):
  original_n = n
  while n != 1 and n != 89:
    if n in results:
      n = results[n]
    else:
      n = sum(x * x for x in map(int, str(n)))

  results[original_n] = n
  return n

count = 0
for n in range(1, 10000000):
  if solve(n) == 89:
    count += 1
  if n % 100000 == 0:
    print count, '/', n
print count
