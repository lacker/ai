// The "database" is just a single object in memory.

let database = {}

export function all(tableName) {
  if (!database[tableName]) {
    database[tableName] = []
  }
  return database[tableName]
}

export function push(tableName, data, limit) {
  let table = all(tableName)
  table.push(data)
  if (limit) {
    while (table.length > limit) {
      table.shift()
    }
  }
}
