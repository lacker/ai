#!/usr/bin/env python

'Yields all proper divisors, out of order'
def proper_divisors(n):
  yield 1
  d = 2
  while d * d <= n:
    if n % d == 0:
      yield d
      other = n / d
      if other != d:
        yield other
    d += 1

def is_abundant(n):
  return sum(proper_divisors(n)) > n
    
small_numbers = range(1, 28124)
    
abundants = [x for x in small_numbers if is_abundant(x)]

is_sum = set()
for a in abundants:
  for b in abundants:
    if a + b > 28124:
      break
    is_sum.add(a+b)

cannot = [x for x in small_numbers if x not in is_sum]
print cannot
print sum(cannot)
