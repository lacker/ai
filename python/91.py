#!/usr/bin/env python

'''
Whether there is a lattice right triangle with
(0, 0)
(x, y)
(top_x, top_y)
and (x, y) is the right angle point.
Needs top_x >= 0
'''
def has_lattice_right(x, y, top_y):
  assert top_y > y

  # Must satisfy:
  # (top_x - x, top_y - y) ~ (-y, x)
  # So solve for top_x:
  # (top_x - x) / -y = (top_y - y) / x
  # top_x = y * (y - top_y) / x + x

  top_x = (float(y) * (y - top_y) / x) + x
  return top_x >= 0 and int(top_x) == top_x

  
'''
Returns how many right triangles have coordinates
(0, 0)
(x, _y)
(_x, y)
where _x <= x and _y <= y
'''
def precise(x, y):
  # You have one where _x = 0, _y = 0
  # Then you have two with a point at (x, y)
  # For three where the right angle itself is grid-aligned
  answer = 3

  # Now you need to check for triangles where the right angle
  # is non-grid-aligned.
  for _y in range(1, y):
    if has_lattice_right(x, _y, y):
      answer += 1
  for _x in range(1, x):
    if has_lattice_right(y, _x, x):
      answer += 1

  return answer

CACHE = {}
CACHE[(1, 1)] = 3

def imprecise(x, y):
  key = (x, y)
  if key in CACHE:
    return CACHE[key]

  if x == 0 or y == 0:
    return 0

  CACHE[key] = (-imprecise(x - 1, y - 1)
                + imprecise(x, y - 1)
                + imprecise(x - 1, y)
                + precise(x, y))
  return CACHE[key]

print imprecise(50, 50)


