/**
 * @param {number} N
 * @param {number} K
 * @return {number[]}
 */
var numsSameConsecDiff = function(N, K) {
  const next = (signs) => {
    let i = signs.length - 1
    while (signs[i]) {
      signs[i] = false
      i--
    }

    signs[i] = true

    return signs
  }
  const toNumber = (dgts) => {
    let r = 0
    for (let i = 0; i < N; i++) {
      r *= 10
      r += dgts[i]
    }

    return r
  }

  const ans = []

  let digits = []
  let signs = []
  for (let i = 0; i < N; i++) {
    digits.push(0)
    signs.push(false)
  }

  if (N == 1) {
    return ans
  }
  if (K == 0) {
    for (let i = 1; i < 10; i++) {
      for (let j = 0; j < N; j++) {
        digits[j] = i
      }
      ans.push(toNumber(digits))
    }

    return ans
  }

  for (let i = 1; i < 10; i++) {
    digits[0] = i
    for (let j = 1; j < N; j++) {
      digits[j] = 0
    }

    let saved = [ ...digits ]

    while (!signs[0]) {
      digits = [ ...saved ]
      let valid = true
      for (let j = 1; valid && j < digits.length; j++) {
        let next = digits[j - 1] + (signs[j] ? 1 : -1) * K
        valid = next >= 0 && next < 10
        digits[j] = next
      }
      //console.log(digits)
      if (valid) {
        ans.push(toNumber(digits))
      }

      signs = next(signs)
    }

    signs[0] = false
  }

  return ans
}

console.log(numsSameConsecDiff(3, 7), [ 181, 292, 707, 818, 929 ])
console.log(numsSameConsecDiff(2, 1))
console.log([ 10, 12, 21, 23, 32, 34, 43, 45, 54, 56, 65, 67, 76, 78, 87, 89, 98 ])
