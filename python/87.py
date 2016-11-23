#!/usr/bin/env python

def is_prime(n):
  if n < 2:
    return False
  if n == 2:
    return True
  if n % 2 == 0:
    return False
  d = 3
  while True:
    if d * d > n:
      return True
    if n % d == 0:
      return False
    d += 2


limit = 50000000
squares = []
cubes = []
forks = []

for p in range(1, limit):
  if not is_prime(p):
    continue
    
  square = p * p
  if square > limit:
    break
  squares.append(square)
  
  cube = p * p * p
  if cube > limit:
    continue
  cubes.append(cube)
    
  fork = p * p * p * p
  if fork > limit:
    continue
  forks.append(fork)

print len(squares), 'squares'
print len(cubes), 'cubes'
print len(forks), 'forks'

sums = set()
for fork in forks:
  for cube in cubes:
    if fork + cube > limit:
      break
      
    for square in squares:
      s = fork + cube + square
      if s > limit:
        break
      sums.add(s)

print len(sums), 'sums'
