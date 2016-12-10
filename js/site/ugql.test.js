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
  expect(sum(1, 2)).toBe(3);
});
