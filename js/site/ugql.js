const assert = require('assert')
const { graphql, parse } = require('graphql')

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

  // TODO: handle arguments

  let result = {}
  let promises = []
  if (query.selectionSet) {
    for (let field of query.selectionSet.selections) {
      assert.equal(field.kind, 'Field')
      let key = field.name.value;
      if (field.selectionSet) {
        promises.push(run(data[key], field).then(r => {
          result[key] = r;
        }))
      } else {
        result[key] = data[key]
      }
    }
  }

  return Promise.all(promises).then(() => result);
}

module.exports = { run }
