#!/usr/bin/env python

'''
Number of rectangles in a x * y rectangle
'''
MEMO = {}
def numrect(x, y):
  if x < y:
    return numrect(y, x)
  if y <= 0:
    return 0
  key = (x, y)
  if key in MEMO:
    return MEMO[key]

  answer = (numrect(x - 1, y)
            + numrect(x, y - 1)
            - numrect(x - 1, y - 1)
            + x * y)
  MEMO[key] = answer
  return answer

def go():
  target = 2000000
  best_distance = target
  best_tuple = None, None
  for x in range(2, target):
    for y in range(1, x):
      num = numrect(x, y)
      dist = abs(num - target)
      if dist < best_distance:
        best_distance = dist
        best_tuple = (x, y)
        print 'numrect(%d, %d) = %d' % (x, y, num)
      if num > target:
        if y == 1:
          return
        break

go()
