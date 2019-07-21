/**
 * @param {number[]} cells
 * @param {number} N
 * @return {number[]}
 */
var prisonAfterNDays = function(cells, N) {
  const seen = new Map()
  const nextDay = (cells) => {
    const tmp = [ ...cells ]
    for (let i = 0; i < cells.length; i++) {
      tmp[i] = 1 * (i > 0 && i < cells.length - 1 && cells[i - 1] === cells[i + 1])
    }

    return tmp
  }

  for (let d = 1; d <= N; d++) {
    cells = nextDay(cells)
    const key = cells.join(',')
    //console.log(d, key)
    if (seen.has(key)) {
      const loopLen = d - seen.get(key)
      d += loopLen * Math.floor((N - d) / loopLen)
    } else {
      seen.set(key, d)
      //d++
    }
  }

  return cells
}

console.log(prisonAfterNDays([ 0, 1, 0, 1, 1, 0, 0, 1 ], 7))
console.log(prisonAfterNDays([ 1, 0, 0, 1, 0, 0, 1, 0 ], 1000000000))
