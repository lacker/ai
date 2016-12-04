#!/usr/bin/env python

months = [
  'NaM',
  'Jan',
  'Feb',
  'Mar',
  'Apr',
  'May',
  'Jun',
  'Jul',
  'Aug',
  'Sep',
  'Oct',
  'Nov',
  'Dec',
]

dows = [
  'Sun',
  'Mon',
  'Tue',
  'Wed',
  'Thu',
  'Fri',
  'Sat',
]

'''
A leap year occurs on any year evenly divisible by 4,
but not on a century unless it is divisible by 400.
'''
def is_leap_year(n):
  if n % 4 != 0:
    return False
  if n % 400 == 0:
    return True
  if n % 100 == 0:
    return False
  return True

def num_days(year, month):
  if month == 2:
    if is_leap_year(year):
      return 29
    return 28
  if month in [9, 4, 6, 11]:
    return 30
  return 31

def next_day(year, month, day):
  day = day + 1
  if day > num_days(year, month):
    if month == 12:
      return year + 1, 1, 1
    return year, month + 1, 1
  return year, month, day

year, month, day = 1900, 1, 1

dow = 1 # Monday

# Num Sundays on the first of the month
count = 0
foms = 0

while (year, month, day) != (2000, 12, 31):
  if day == 1 and year >= 1901:
    print year, months[month], day, '=', dows[dow]
    foms += 1
    if dow == 0:
      count += 1

  dow = (dow + 1) % 7
  year, month, day = next_day(year, month, day)
  
print 'answer:', count
print 'foms:', foms
