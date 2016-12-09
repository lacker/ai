// The "database" is just a single object in memory.

let database = {}

function all(tableName) {
  if (!database[tableName]) {
    database[tableName] = []
  }
  return database[tableName]
}

function push(tableName, data, limit) {
  if (!tableName) {
    throw Error('push needs a tableName')
  }
  console.log('pushing:', data)
  let table = all(tableName)
  table.push(data)
  if (limit) {
    while (table.length > limit) {
      table.shift()
    }
  }
}

export default { all, push }
