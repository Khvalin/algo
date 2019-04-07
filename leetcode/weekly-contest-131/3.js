/**
 * @param {string[]} queries
 * @param {string} pattern
 * @return {boolean[]}
 */
var camelMatch = function(queries, pattern) {
  const test = (q) => {
    let i = 0

    for (const ch of q) {
      if (ch === pattern[i]) {
        i++
        continue
      }

      if (ch.toUpperCase() === ch) {
        return false
      }
    }

    return i === pattern.length
  }

  return queries.map((q) => test(q))
}

console.log(camelMatch([ 'FooBar', 'FooBarTest', 'FootBall', 'FrameBuffer', 'ForceFeedBack' ], 'FB'))
console.log(camelMatch([ 'FooBar', 'FooBarTest', 'FootBall', 'FrameBuffer', 'ForceFeedBack' ], 'FoBaT'))
