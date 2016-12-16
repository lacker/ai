const assert = require('assert')
const { graphql, parse } = require('graphql')

// Resolves a graphql value, which is basically representing a literal.
function resolveValue(value) {
  if (value.kind == 'IntValue') {
    return parseInt(value.value);
  }

  throw new Error('this code needs to handle value.kind =', value.kind)
}

// TODO: refactor out the field resolver

function run(data, query) {
  // Resolve promises
  if (data.then) {
    return data.then(d => run(d, query))
  }

  // Find the actual graphql
  if (typeof query == 'string') {
    let document = parse(query)
    assert.equal(document.kind, 'Document')
    query = document.definitions[0]
    assert.equal(query.kind, 'OperationDefinition')
  }

  let result = {}
  let promises = []

  assert(query.selectionSet)
  for (let field of query.selectionSet.selections) {
    assert.equal(field.kind, 'Field')
    let key = field.name.value;
    if (field.selectionSet) {
      promises.push(run(data[key], field).then(r => {
        result[key] = r;
      }))
      continue
    }

    if (field.arguments && field.arguments.length > 0) {
      let args = {}
      for (let arg of field.arguments) {
        args[arg.name.value] = resolveValue(arg.value)
      }
      result[key] = data[key](args)
      continue
    }

    if (data[key].then) {
      promises.push(data[key].then(val => {
        result[key] = val;
      }))
      continue
    }
    result[key] = data[key]
  }

  return Promise.all(promises).then(() => result);
}

module.exports = { run }
