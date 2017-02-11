
// Tries to remove the items in b from a.
// If the items in b are not in a, returns null.
// Otherwise returns a list of the other items.
function diff(a, b) {
  let bIndex = 0;
  let answer = [];
  for (let item of a) {
    if (b[bIndex] === item) {
      bIndex++;
    } else {
      answer.push(item);
    }
  }
  if (answer.length == a.length - b.length) {
    return answer;
  } else {
    return null;
  }
}

function shuffle(arr) {
  for (let i = arr.length - 1; i > 0; i--) {
    let j = Math.floor(Math.random() * (i + 1));
    let tmp = arr[i];
    arr[i] = arr[j];
    arr[j] = tmp;
  }
}

function choose(arr) {
  return arr[Math.floor(Math.random() * arr.length)];
}

function logSquare(arr) {
  let line = [];
  for (let value of arr) {
    line.push(value);
    if (line.length == size) {
      console.log(line);
      line = [];
    }
  }
}

// Merges two ascending lists, deduping.
function merge(a, b) {
  let answer = [];
  let aIndex = 0;
  let bIndex = 0;
  while (aIndex < a.length || bIndex < b.length) {
    if (aIndex >= a.length) {
      answer.push(b[bIndex]);
      bIndex++;
    } else if (bIndex >= b.length) {
      answer.push(a[aIndex]);
      aIndex++;
    } else if (a[aIndex] < b[bIndex]) {
      answer.push(a[aIndex]);
      aIndex++;
    } else if (a[aIndex] > b[bIndex]) {
      answer.push(b[bIndex]);
      bIndex++;
    } else {
      // Dupe
      answer.push(a[aIndex]);
      aIndex++;
      bIndex++;
    }
  }
  return answer;
}

// A "cage" here is a list of variable indices from 0 .. sideLength ^ 2 - 1
// Returns a list of cages
function randomCages(sideLength) {
  // indices is just the order we'll process the indices in
  let indices = [];

  // cageForIndex maps index to the cage it's in, or null if not known yet
  let cageForIndex = [];

  for (let i = 0; i < sideLength * sideLength; i++) {
    indices.push(i);
    cageForIndex.push(null);
  }
  shuffle(indices);

  let cages = [];
  for (let index of indices) {
    let adjacent = [];
    if (index % sideLength !== 0) {
      adjacent.push(index - 1);
    }
    if ((index + 1) % sideLength !== 0) {
      adjacent.push(index + 1);
    }
    if (index - sideLength >= 0) {
      adjacent.push(index - sideLength);
    }
    if (index + sideLength < sideLength * sideLength) {
      adjacent.push(index + sideLength);
    }

    // Lower cage score is better
    let bestCage = null;
    let bestCageScore = 100;

    for (let index of adjacent) {
      let cage = cageForIndex[index];
      if (cage === null) {
        continue;
      }

      let score = cages[cage].length;
      if (score >= 4) {
        continue;
      }
      if (score < bestCageScore) {
        bestCage = cage;
        bestCageScore = score;
      }
    }

    if (bestCage === null) {
      bestCage = cages.length;
      cages.push([]);
    }
    cageForIndex[index] = bestCage;
    cages[bestCage].push(index);
  }

  return cages;
}

// Returns a list of lists.
// Subsets are in the same order as the superset
function allSubsets(items, numItems) {
  if (numItems === 0) {
    return [[]];
  }
  if (numItems > items.length) {
    return [];
  }
  let answer = [];

  // First handle the cases where we do take the first item
  let firstItem = items[0];
  let otherItems = items.slice(1);
  for (let tail of allSubsets(otherItems, numItems - 1)) {
    answer.push([firstItem].concat(tail));
  }

  // Then handle the cases where we don't take the first item
  return answer.concat(allSubsets(otherItems, numItems));
}

// Makes the containers for a particular cage
// operation can be either '*' or '+'
// result is what everything is supposed to go into
// each container should be numValues values in [1, size]
function makeContainers(operation, result, numValues, size) {
  let containers = [];
  let domain = [];
  for (let i = 1; i <= size; i++) {
    domain.push(i);
  }
  for (let container of allSubsets(domain, numValues)) {
    let res;
    switch (operation) {
      case '+':
      res = container.reduce((a, b) => (a + b), 0);
      break;
      case '*':
      res = container.reduce((a, b) => (a * b), 1);
      break;
      default:
      throw new Error('bad operation: ' + operation);
    }
    if (res === result) {
      containers.push(container);
    }
  }
  return containers;
}

// cage is a list of indices in values.
// it is thus the "variables" arg to addConstraint.
// returns an object with {description, containers}.
function makeCageConstraint(values, cage) {
  // XXX
}

// Intersects two ascending lists.
function intersect(a, b) {
  let answer = [];
  let aIndex = 0;
  let bIndex = 0;
  while (aIndex < a.length && bIndex < b.length) {
    if (a[aIndex] < b[bIndex]) {
      aIndex++;
    } else if (a[aIndex] > b[bIndex]) {
      bIndex++;
    } else {
      // They must be equal
      answer.push(a[aIndex]);
      aIndex++;
      bIndex++;
    }
  }
  return answer;
}

// soFar is an ascending list of numbers
// containers is a list of ascending lists of numbers
// This returns an ascending list of all numbers that could be added to
// soFar while keeping soFar as a sublist of one of the containers.
function possibilities(soFar, containers) {
  let answer = [];
  for (let container of containers) {
    let d = diff(container, soFar);
    if (d !== null) {
      answer = merge(answer, d);
    }
  }
  return answer;
}

// Does backtracking
class Puzzle {
  constructor(numVariables) {
    this.variables = Array(numVariables).fill(null);

    // Each constraint is an object with:
    // variables: a list of ints, indices in this.variables. In order
    // containers:
    //       A list of sets. The constraint is that the variables must
    //       map to one of these sets. A "set" here is an ascending
    //       list of integers.
    // description: a string describing this constraint
    this.constraints = [];

    // Maps to a list of indices in this.constraints
    this.constraintsForVariable = [];
    for (let v of this.variables) {
      this.constraintsForVariable.push([]);
    }
  }

  // The constraint is that the variables specified in 'variables' must
  // be a subset of one of the lists in 'containers'.
  addConstraint(variables, containers, description) {
    let index = this.constraints.length;
    this.constraints.push({
      variables: variables,
      containers: containers,
      description: description,
    });
    for (let v of variables) {
      this.constraintsForVariable[v].push(index);
    }
  }

  // Returns a list of the possible values that could come next.
  possibleNext(values) {
    if (values.length >= this.variables.length) {
      throw 'values is too long for possibleNext';
    }
    if (this.variables[values.length] !== null) {
      return [this.variables[values.length]];
    }

    // The constraints that are relevant to the next value
    let constraintIndices = this.constraintsForVariable[values.length];

    // If answer is non-null, it's a superset of the possible values.
    // This is because any possible value must meet each constraint.
    let answer = null;

    for (let constraintIndex of constraintIndices) {
      let constraint = this.constraints[constraintIndex];

      // Let's find partial solutions, that are at least ok with
      // this constraint.
      // First figure out what values are already filled in, for this
      // constraint.
      let alreadyFilled = [];
      for (let index of constraint.variables) {
        if (index >= values.length) {
          break;
        }
        alreadyFilled.push(values[index]);
      }
      alreadyFilled.sort(); // NOTE: this assumes numbers are < 10 !

      let partials = possibilities(alreadyFilled, constraint.containers);
      if (answer === null) {
        answer = partials;
      } else {
        answer = intersect(answer, partials);
      }

      // Shortcut
      if (answer.length == 0) {
        return answer;
      }
    }

    return answer;
  }

  // Solves with backtracking.
  // values is the variable values that have been figured out so far.
  // method can be: 'reverse' or 'random'. others do it in order
  // Returns a list of values if there's a solution.
  // Returns null otherwise.
  solve(values, method) {
    if (values.length === this.variables.length) {
      return values;
    }
    let possible = this.possibleNext(values);
    if (method === 'reverse') {
      possible.reverse();
    } else if (method === 'random') {
      shuffle(possible);
    }
    for (let nextValue of possible) {
      const answer = this.solve(values.concat([nextValue]), method);
      if (answer !== null) {
        return answer;
      }
    }
    return null;
  }
}

// Creates a Puzzle whose constraints just represent a valid Sudoku board.
// 'size' is the length of one side length of the square.
function anySudoku(size) {
  let puzzle = new Puzzle(size * size);

  // Each row and column has a container list with just one
  // legitimate container - the list of 1..size numbers
  let validNumbers = [];
  for (let i = 1; i <= size; i++) {
    validNumbers.push(i);
  }
  const containers = [validNumbers];

  for (let i = 0; i < size; i++) {
    let row = [];
    let col = [];
    for (let j = 0; j < size; j++) {
      row.push(i * size + j);
      col.push(j * size + i);
    }
    puzzle.addConstraint(row, containers);
    puzzle.addConstraint(col, containers);
  }

  return puzzle;
}



const size = 6;

// This code is just to print cages reasonably.
let cages = randomCages(6);
let x = [];
for (let i = 0; i < cages.length; i++) {
  let cage = cages[i];
  for (let index of cage) {
    x[index] = i;
  }
}
logSquare(x);

console.log(makeContainers('*', 30, 3, 6));
