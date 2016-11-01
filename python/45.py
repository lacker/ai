#!/usr/bin/env python

tri = set()
penta = set()
hexa = set()

n = 0
while True:
  n += 1

  t = n * (n + 1) / 2
  p = n * (3 * n - 1) / 2
  h = n * (2 * n - 1)

  tri.add(t)
  penta.add(p)
  hexa.add(h)

  for num in (t, p, h):
    if num in tri and num in penta and num in hexa:
      print num
