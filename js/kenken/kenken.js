
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
    // sets: A list of sets. The constraint is that the variables must
    //       map to one of these sets. A "set" here is an ascending
    //       list of integers.
    this.constraints = [];

    // Maps to a list of indices in this.constraints
    this.constraintsForVariable = [];
    for (let v of this.variables) {
      this.constraintsForVariable.push([]);
    }
  }

  addConstraint(variables, sets) {
    let index = this.constraints.length;
    this.constraints.push({
      variables: variables,
      sets: sets,
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

    // The constraints that are relevant to the next value
    let constraints = this.constraintsForVariable[values.length];

    for (let constraint of constraints) {
      // XXX
    }
  }

  // Solves with backtracking.
  // values is the variable values that have been figured out so far.
  solve(values) {
    // XXX
  }
}
