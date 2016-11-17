#!/usr/bin/env python

def digifac(n):
  if n >= 10 or n < 0:
    raise 'fail'
  if n < 2:
    return 1
  return reduce(lambda x, y: x*y, range(1, n + 1))

def sofac(n):
  return sum(digifac(int(ch)) for ch in str(n))

def chainlen(n):
  term = n
  seen = set([term])
  while True:
    term = sofac(term)
    if term in seen:
      return len(seen)
    seen.add(term)

def go():
  answer = 0
  for x in range(1, 1000000):
    c = chainlen(x)
    if c >= 50:
      print 'chainlen(%d) = %d' % (x, c)
    if c == 60:
      answer += 1
    if c > 60:
      raise 'fail'
  print
  print answer

go()


