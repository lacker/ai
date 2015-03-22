#!/usr/bin/python
"""
An interpreter for boxcode.
"""

"""
Parses a line into an S-expression or atom.
S-expressions are just represented as Python lists.

Valid atoms are:
symbols (no whitespace or (), may not start with a number or .)
numbers (anything starting with a number of . is assumed to be a number)

symbols are just represented as strings.
numbers are represented as Python numbers.
"""
def parse(serialized):
  pass
  

