#!/usr/bin/env python

from math import sqrt

billion = 1000000000

answer = 0

# k has to be odd

# Check k, k, k + u
for k in range(3, billion, 2):
  for u in (-1, +1):
    perimeter = 3 * k + u
    if perimeter > billion:
      break
    semi = perimeter / 2
    area = sqrt(semi * (semi - k) * (semi - k) * (semi - k - u))
    if int(area) == area:
      print k, k, k + u
      answer += perimeter
  
'''
So semi * (semi - k - u) must be square
== (3k + u) / 2 * (k - u) / 2 is square
(3k + u) (k - u) / 4 is square
'''

print answer
