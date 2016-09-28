#!/usr/bin/env python

CACHE = {}

def get_key(total, coins):
  return ','.join(map(str, [total] + coins))

def set_cache(total, coins, answer):
  print 'setting', total, '::', coins, '=', answer
  CACHE[get_key(total, coins)] = answer

def get_cache(total, coins):
  return CACHE.get(get_key(total, coins), None)

def ways_to_make(total, coins):
  print 'calling wtm', total, coins
  if total < 0:
    raise 'total should not be < 0'
    
  if not coins:
    if total == 0:
      return 1
    return 0

  cache = get_cache(total, coins)
  if cache is not None:
    return cache
    
  first_coin = coins[0]
  other_coins = coins[1:]
  answer = 0
  num_first_coins = 0
  while True:
    amount_in_first_coins = first_coin * num_first_coins
    if amount_in_first_coins > total:
      break
    answer += ways_to_make(total - amount_in_first_coins, other_coins)
    num_first_coins += 1

  set_cache(total, coins, answer)
  return answer

print ways_to_make(17, [10, 5, 1])
  
print ways_to_make(200, list(reversed([1, 2, 5, 10, 20, 50, 100, 200])))
