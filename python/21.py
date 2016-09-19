#!/usr/bin/python
def sum_of_divisors(n):
  answer = 0
  for i in range(1, n / 2 + 1):
    if n % i == 0:
      answer += i
  return answer

dmap = {}

def is_amicable(n):
  friend = sum_of_divisors(n)
  if friend == n:
    return False
  double_friend = sum_of_divisors(friend)
  return double_friend == n

answer = 0
for n in range(1, 10000):
  if is_amicable(n):
    print 'amicable:', n
    answer += n
    
print 'answer:', answer
