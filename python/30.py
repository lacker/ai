#!/usr/bin/env python
# The number can be at most 6 digits

def sum_of_fifth_powers(n):
  answer = 0
  for char in str(n):
    answer += int(char) ** 5
  return answer

answer = 0
for x in range(10, 1000000):
  if sum_of_fifth_powers(x) == x:
    print x
    answer += x
print 'answer:', answer
