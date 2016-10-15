#!/usr/bin/env python

# Find the nth pentagonal number
def pentagon(n):
  return (3 * n * n - n) / 2

'''
A good tuple is (a, b, c) where pentagon(a) + pentagon(b) = pentagon(c).

We are looking for pairs of good tuples: (a, b, c) and (b, c, d) are both good.
And we want the min a.
'''

NUM = 10000

nums = [pentagon(n) for n in range(1, NUM + 1)]
pentagons = set(nums)

for a in nums:
  for b in nums:
    if (a - b) in pentagons and (a + b) in pentagons:
      print a, b, a - b, a + b
