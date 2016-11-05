#!/usr/bin/env python

MEMO = {}

def choose(n, k):
  if n < 1:
    raise Exception('n should be >= 1')
  if k > n:
    raise Exception('k should be <= n')
  if n == 1:
    return 1
  if n == k:
    return 1
  if k == 0:
    return 1
    
  key = (n, k)
  if key in MEMO:
    return MEMO[key]

  answer = choose(n - 1, k) + choose(n - 1, k - 1)
  MEMO[key] = answer
  return answer

count = 0
for n in range(1, 101):
  for k in range(n + 1):
    if choose(n, k) > 1000000:
      count += 1
print count
