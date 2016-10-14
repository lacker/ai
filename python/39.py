#!/usr/bin/env python
from math import sqrt

# perimeter -> count of sol'ns
count = {}

for a in range(1, 1000):
  print 'a =', a
  for b in range(a, 1000):
    left_sum = a * a + b * b
    c = int(sqrt(left_sum))
    perimeter = a + b + c
    if perimeter > 1000:
      break
    if c * c != left_sum:
      continue
    count[perimeter] = count.get(perimeter, 0) + 1

print max(count, key=count.get)
