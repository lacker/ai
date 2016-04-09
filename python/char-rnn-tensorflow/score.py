#!/usr/bin/env python3
'''
Scores an input that looks like

>>> # python statement
evaled result

So a success might be

>>> 3 + 5
8

and a failure might be

>>> 5 > 3
False

Then prints some stats.
'''

def main():
  with open('save/output.txt') as f:
    correct = 0
    incorrect = 0
    last_line = ''
    graded = []
    for line in f.readlines():
      line = line.strip()
      graded.append(line)
      
      if last_line.startswith('>>> '):
        code = last_line[4:]

        response = ''
        try:
          response = repr(eval(code))
        except:
          response = 'Error'
          
        if line == response:
          correct += 1
          graded.append('Correct')
        else:
          incorrect += 1
          graded.append('Incorrect')

      last_line = line

    print('%d / %d = %.2f%% correct\n' % (
      correct, correct + incorrect,
      100.0 * correct / (correct + incorrect)))

    for line in graded:
      print(line)
          

if __name__ == '__main__':
  main()
