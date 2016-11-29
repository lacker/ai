#!/usr/bin/env python

from math import sqrt

billion = 1000000000

answer = 0

# k has to be odd

# Check k, k, k + u
for k in xrange(3, billion, 2):
  for u in (-1, +1):
    perimeter = 3 * k + u
    if perimeter > billion:
      print 'answer:', answer
      raise Exception('done')
    semi = perimeter / 2
    # z = area / (semi - k) - heron's formula
    sq = semi * (semi - k - u)
    z = int(sqrt(sq))
    if z * z == sq:
      print k, k, k + u
      answer += perimeter
