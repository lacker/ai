#!/usr/bin/env python

def revnum(n):
  return int(''.join(list(reversed(str(n)))))

def isly(orig_n):
  n = orig_n
  for i in range(60):
    r = revnum(n)
    if r == n and i > 0:
      return False
    n += r
  print orig_n
  return True

print len(filter(isly, range(10000)))
