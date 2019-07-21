/**
 * @param {number[]} arr
 * @return {number}
 */
var mctFromLeafValues = function(arr) {
  const even = (arr) => {
    let level = [ ...arr ]
    let res = 0
    while (level.length > 1) {
      console.log(level)
      for (let i = 0; i < level.length; i += 2) {
        level[i] *= level[i + 1]
        res += level[i]
      }

      level = level.filter((val, i) => i % 2 === 0)
    }

    return res + level[0]
  }

  if (arr.length % 2 === 0) {
    return even(arr)
  }

  const copy = [ ...arr ]

  const left = copy.shift()
  const right = arr.pop()
  return Math.min(left + even(copy), right + even(arr))
}

console.log(mctFromLeafValues([ 6, 2, 4 ]))
