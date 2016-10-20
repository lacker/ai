#!/usr/bin/env python

# Returns the cycle length of 1 / d
def cycle_len(d):
  while d % 2 == 0:
    d = d / 2
  while d % 5 == 0:
    d = d / 5
  if d < 2:
    return 0
  answer = 1
  while True:
    denom = int('9' * answer)
    if denom % d == 0:
      return answer
    answer += 1

ans = 1
for x in range(1, 1000):
  ans = max(ans, cycle_len(x))
  print x, 'has cycle length', cycle_len(x)
print ans
