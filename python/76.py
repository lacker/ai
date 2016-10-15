#!/usr/bin/env python

memo = {}

# From largest to smallest
def num_sums(n, max_kid):
  if n == 0:
    return 1

  if max_kid > n:
    return num_sums(n, n)
    
  key = (n, max_kid)
  if key in memo:
    return memo[key]

  answer = 0
  for first_kid in range(1, max_kid + 1):
    answer += num_sums(n - first_kid, first_kid)
  memo[key] = answer
  return answer

print 1, num_sums(1, 1) - 1
print 2, num_sums(2, 2) - 1
print 3, num_sums(3, 3) - 1
print 4, num_sums(4, 4) - 1
print 5, num_sums(5, 5) - 1
print 100, num_sums(100, 100) - 1
  
