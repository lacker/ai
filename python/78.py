#!/usr/bin/env python

'-1 ** k'
def negabase(k):
  if k % 2 == 0:
    return 1
  else:
    return -1

'''
The number of partitions of n.
'''
MEMO = {}
def pcount(n):
  if n == 0:
    return 1
  if n in MEMO:
    return MEMO[n]
    
  k = 1
  answer = 0
  while True:
    pentagon = k * (3 * k - 1) / 2
    arg = n - pentagon
    if arg < 0:
      MEMO[n] = answer
      return MEMO[n]

    answer += negabase(k - 1) * pcount(arg)
    if k > 0:
      k = -k
    else:
      k = -k + 1
    
for n in range(10000000):
  p = pcount(n)
  print 'pcount(%d) = %d' % (n, p)

  if p % 1000000 == 0:
    print 'DONE'
    break
