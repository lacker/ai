const { run } = require('./ugql')

test('nested queries', () => {
  let data = {
    foo: 1,
    bar: {
      baz: 2
    }
  }
  let query = `{
    foo
    bar {
      baz
    }
  }`
  return run(data, query).then(output => {
    expect(output).toEqual(data)
    expect(output).not.toBe(data)
  })
})

test('nested promise', () => {
  let data = {
    foo: Promise.resolve({
      bar: 3
    })
  }
  let query = `{
    foo {
      bar
    }
  }`
  return run(data, query).then(output => {
    expect(output).toEqual({ foo: { bar: 3 }})
  })
})

test('promise of value', () => {
  let data = {
    foo: Promise.resolve(2)
  }
  let query = `{
    foo
  }`
  return run(data, query).then(output => [
    expect(output).toEqual({ foo: 2 })
  ])
})

test('arguments', () => {
  let data = {
    addOne: ({x}) => (x + 1)
  }
  let query = `{
    addOne(x: 2)
  }`
  return run(data, query).then(output => {
    expect(output).toEqual({ addOne: 3 })
  })
})

test('arguments then nesting', () => {
  let data = {
    foo: ({x}) => ({ bar: x })
  }
  let query = `{
    foo(x: 1) {
      bar
    }
  }`
  return run(data, query).then(output => {
    expect(output).toEqual({ foo: { bar: 1 }})
  })
})

test('zero-argument functions', () => {
  let data = {
    foo: () => 2
  }
  let query = `{
    foo
  }`
  return run(data, query).then(output => {
    expect(output).toEqual({ foo: 2 })
  })
})

test('string responses', () => {
  let data = {
    foo: 'foo'
  }
  let query = `{
    foo
  }`
  return run(data, query).then(output => {
    expect(output).toEqual({ foo: 'foo' })
  })
})

test('string arguments', () => {
  let data = {
    double: ({x}) => (x + x)
  }
  let query = `{
    double(x: "foo")
  }`
  return run(data, query).then(output => {
    expect(output).toEqual({ double: 'foofoo' })
  })
})

test('array data', () => {
  let data = {
    foo: [{x: 1, y: 0}, {x: 2}]
  }
  let query = `{
    foo {
      x
    }
  }`
  return run(data, query).then(output => {
    expect(output).toEqual({ foo: [{x: 1}, {x: 2}]})
  })
})
