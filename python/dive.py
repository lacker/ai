#!/usr/bin/env python

def evaluate(alist):
  '''
  Evalutes a list of things.
  ['quote' foo] evaluates to foo
  TODO: more stuff
  '''
  if not alist:
    raise Exception('cannot evaluate the empty list')
    
  first = alist[0]
  rest = alist[1:]

  if first == 'quote':
    if len(rest) != 1:
      raise Exception('quote takes exactly one arg')
    return rest[0]
  else:
    raise Exception('unrecognized functiony thing: ' + first)
