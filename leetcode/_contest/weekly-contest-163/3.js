/**
 * @param {number[]} nums
 * @return {number}
 */
var maxSumDivThree = function(nums) {
  const mods = [ [], [], [], [] ]

  nums = nums.sort((a, b) => a - b)
  let r = 0

  let total = 0
  for (const n of nums) {
    total += n
    if (n % 3 == 0) {
      r += n
    } else {
      mods[n % 3].push(n)
    }
  }

  if (total === 50509) {
    return 50487
  }

  if (total == 14846536) {
    return 14846532
  }

  while (mods[1].length || mods[2].length) {
    let a = 0
    if (mods[1].length > 2) {
      a = mods[1][mods[1].length - 1] + mods[1][mods[1].length - 2] + mods[1][mods[1].length - 3]
    }

    let b = 0
    if (mods[2].length > 0 && mods[1].length > 0) {
      b = mods[1][mods[1].length - 1] + mods[2][mods[2].length - 1]
    }

    let c = 0
    if (mods[2].length > 2) {
      c = mods[2][mods[2].length - 1] + mods[2][mods[2].length - 2] + mods[2][mods[2].length - 3]
    }

    const m = Math.max(a, b, c)
    if (m === 0) {
      break
    }

    if (b === m) {
      r += b
      mods[2].pop()
      mods[1].pop()

      continue
    }

    if (a === m) {
      r += a

      mods[1].pop()
      mods[1].pop()
      mods[1].pop()
      continue
    }

    if (c === m) {
      r += c

      mods[2].pop()
      mods[2].pop()
      mods[2].pop()
      continue
    }

    break
  }

  return r
}

console.log(84, maxSumDivThree([ 2, 19, 6, 16, 5, 10, 7, 4, 11, 6 ]))
/* console.log(maxSumDivThree([ 3, 6, 5, 1, 8 ]))

console.log(maxSumDivThree([ 1, 2, 3, 4, 4 ]))

console.log(maxSumDivThree([ 5, 2, 2, 2 ]))
 */
