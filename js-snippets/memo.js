

/**
 * @param {Function} func
 * @param {(args:[]) => string }  [resolver] - cache key generator
 */
function memo(func, resolver) {
  if (!resolver) {
    resolver = (...args) => JSON.stringify(args)
  }

  const m = new Map()
  return function(...args) {
    const key = resolver(...args)
    if (m.has(key)) {
      return m.get(key)
    }

    const res = func.apply(this, args)
    m.set(key, res)

    return res
  }
}

// TEST
let callCount = 0
const func = (a, b) => {
  callCount += 1
  return a + b
}
const memoed = memo(func, (a, b) => (a + b) % 2 === 0 ? 'even' : 'odd')

memoed(1, 2)
memoed(1, 4)
memoed(1, 3)
memoed(11, 31)
