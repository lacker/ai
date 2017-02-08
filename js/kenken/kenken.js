
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

// Intersects two ascending lists.
function intersect(a, b) {
  let answer = [];
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
    this.constraints = [];

    // Maps to a list of indices in this.constraints
    this.constraintsForVariable = [];
    for (let v of this.variables) {
      this.constraintsForVariable.push([]);
    }
  }

  // The constraint is that the variables specified in 'variables' must
  // be a subset of one of the lists in 'containers'.
  addConstraint(variables, containers) {
    let index = this.constraints.length;
    this.constraints.push({
      variables: variables,
      containers: containers,
    });
    for (let v of variables) {
      this.constraintsForVariable[v].push(index);
    }
  }

  // Returns a list of the possible values that could come next.
  possibleNext(values) {
    console.log('XXX possibleNext(', values, ')');
    if (values.length >= this.variables.length) {
      throw 'values is too long for possibleNext';
    }

    // The constraints that are relevant to the next value
    let constraintIndices = this.constraintsForVariable[values.length];

    // If answer is non-null, it's a superset of the possible values.
    // This is because any possible value must meet each constraint.
    let answer = null;

    for (let constraintIndex of constraintIndices) {
      let constraint = this.constraints[constraintIndex];
      console.log('XXX constraint:', constraint);

      // Let's find partial solutions, that are at least ok with
      // this constraint.
      let partials = possibilities(values, constraint.containers);
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
  // Returns a list of values if there's a solution.
  // Returns null otherwise.
  solve(values) {
    let possible = this.possibleNext(values);
    for (let nextValue of possible) {
      const answer = solve(values.concat([nextValue]));
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

const size = 4;
let board = anySudoku(size);
let solution = board.solve([]);
let line = [];
for (let value of solution) {
  line.push(value);
  if (line.length == size) {
    console.log(line);
    line = [];
  }
}
