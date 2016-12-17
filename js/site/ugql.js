const assert = require('assert')
const { graphql, parse } = require('graphql')

// Resolves a graphql value, which is basically representing a literal.
function resolveValue(value) {
  if (value.kind == 'IntValue') {
    return parseInt(value.value);
  }

  throw new Error('this code needs to handle value.kind =', value.kind)
}

function run(data, query) {
  // Resolve promises
  if (data.then) {
    return data.then(d => run(d, query))
  }

  // If we get a string, assume it's graphql with one operation
  if (typeof query == 'string') {
    let document = parse(query)
    assert.equal(document.kind, 'Document')
    assert(document.definitions.length == 1)
    query = document.definitions[0]
    assert.equal(query.kind, 'OperationDefinition')
  }

  // If we get something with arguments, resolve those first
  if (query.arguments && query.arguments.length > 0) {
    assert.equal(typeof data, 'function')
    let args = {}
    for (let arg of query.arguments) {
      args[arg.name.value] = resolveValue(arg.value)
    }
    let q = Object.assign({}, query)
    q.arguments = null
    return run(data(args), q)
  }

  // Functions when data is expected are ok, they get treated
  // like thunks
  if (typeof data == 'function') {
    return run(data(), query)
  }

  if (!query.selectionSet) {
    return Promise.resolve(data)
  }

  // Recurse on the subselections

  let result = {}
  let promises = []

  for (let field of query.selectionSet.selections) {
    assert.equal(field.kind, 'Field')
    let key = field.name.value;
    promises.push(run(data[key], field).then(r => {
      result[key] = r;
    }))
  }

  return Promise.all(promises).then(() => result);
}

module.exports = { run }
