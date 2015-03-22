#!/usr/bin/python
"""
An interpreter for boxcode.
"""

"""
Parses a line into an S-expression or atom.
S-expressions are just represented as Python lists.
See http://norvig.com/lispy.html for inspiration.

Valid atoms are:
symbols (no whitespace or (), may not start with a number or .)
numbers (anything starting with a number of . is assumed to be a number)

symbols are just represented as strings.
numbers are represented as Python numbers.
"""
def parse(serialized):
  tokens = serialized.replace("(", " ( ").replace(")", " ) ").split()

  # Expressions that are partially created
  answer = None
  partials = [[]]

  for token in tokens:
    if answer is not None:
      raise "oops too many tokens"
      
    if token == "(":
      partials.append([])
    elif token == ")":
      expression = partials.pop()
      partials[-1].append(expression)
    else:
      # It's an atom
      if token[0] in ".0123456789":
        partials[-1].append(float(token))
      else:
        partials[-1].append(token)
  
  # partials[0][0] should be size 1 and have the answer
  assert len(partials) == 1
  assert len(partials[0]) == 1
  return partials[0][0]

