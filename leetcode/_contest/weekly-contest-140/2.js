/**
 * @param {string} tiles
 * @return {number}
 */
var numTilePossibilities = function(tiles) {
  const fact = (n) => {
    let r = 1
    for (let i = 2; i <= n; i++) {
      r *= i
    }

    return r
  }

  const countPermutations = (s) => {
    let count = {}
    let set = []

    for (let i = 0; i < s.length; i++) {
      if (!count[s[i]]) {
        set.push(s[i])
      }
      count[s[i]] = count[s[i]] || 0
      count[s[i]]++
    }

    let res = fact(s.length)
    for (let i = 0; i < set.length; i++) {
      res /= fact(count[set[i]])
    }

    return res
  }

  let res = 0
  let used = {}
  let arr = [ tiles ]
  while (arr.length > 0) {
    let s = arr.shift()
    if (s.length === 0) {
      continue
    }

    const c = countPermutations(s)

    res += countPermutations(s)
    for (let i = 0; i < s.length; i++) {
      let next = s.substring(0, i) + s.substring(i + 1, s.length)
      let charts = next.split('').sort((a, b) => a.localeCompare(b))
      next = charts.join('')

      if (used[next]) {
        continue
      }
      used[next] = true

      arr.push(next)
    }
  }

  return res
}

console.log(numTilePossibilities('AAB'))
console.log(numTilePossibilities('AAABBC'))
console.log(numTilePossibilities('CDC'))
