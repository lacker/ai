#!/usr/bin/env python

PRIME_MEMO = {}
def is_prime(n):
  if n < 2:
    return False
  if n == 2:
    return True
  if n % 2 == 0:
    return False
  if n in PRIME_MEMO:
    return PRIME_MEMO[n]
    
  d = 3
  while True:
    if d * d > n:
      PRIME_MEMO[n] = True
      return True
    if n % d == 0:
      PRIME_MEMO[n] = False
      return False
    d += 2


'''
The number of ways that n can be the sum of primes.
"maximum" is the maximum prime allowed
'''
MEMO = {}
def nsop(n, maximum=None):
  if maximum is None or maximum > n:
    maximum = n
    
  key = (n, maximum)
  if key in MEMO:
    return MEMO[key]

  if n == 0:
    return 1
  if n == 1:
    return 0
  if maximum < 2:
    return 0
    
  if not is_prime(maximum):
    MEMO[key] = nsop(n, maximum - 1)
    return MEMO[key]

  MEMO[key] = nsop(n, maximum - 1) + nsop(n - maximum, maximum)
  return MEMO[key]

print nsop(10)
  
for n in range(10000000):
  ns = nsop(n)
  if ns > 5000:
    print 'nsop(%d) = %d' % (n, ns)
    break
  if n % 1000 == 0:
    print 'at', n
