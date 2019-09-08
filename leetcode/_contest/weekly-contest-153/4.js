const memo = {}

/**
 * @param {number[]} arr1
 * @param {number[]} arr2
 * @return {number}
 */
var makeArrayIncreasing = function(arr1, arr2) {
  const calc = (arr1, arr2, prev = Number.NEGATIVE_INFINITY) => {
    const key = prev.toString() + ' -- ' + arr1.length + ' -- ' + arr2.length

    if (memo[key]) {
      return memo[key]
    }

    if (arr1.length === 0) {
      memo[key] = 0
      return 0
    }

    while (arr2.length && arr2[0] <= prev) {
      arr2.shift()
    }

    if (prev >= arr1[0] && !arr2.length) {
      memo[key] = Number.POSITIVE_INFINITY
      return Number.POSITIVE_INFINITY
    }

    let res1 = Number.POSITIVE_INFINITY
    if (arr1[0] > prev) {
      res1 = calc([ ...arr1.slice(1) ], [ ...arr2 ], arr1[0])
    }

    let res2 = Number.POSITIVE_INFINITY
    if (arr2.length) {
      //  console.log(arr2)
      res2 = 1 + calc([ ...arr1.slice(1) ], [ ...arr2.slice(1) ], arr2[0])
    }

    const r = Math.min(res1, res2)
    memo[key] = r

    return r
  }

  arr2 = arr2.sort((a, b) => a - b)

  const res = calc(arr1, arr2)

  return res > arr1.length ? -1 : res
}

let res
/* 
res = makeArrayIncreasing([ 1, 2, 3, 4, 4 ], [ 0, 1, 2, 3 ])
console.log(res)

res = makeArrayIncreasing([ 1, 5, 3, 6, 7 ], [ 1, 3, 2, 4 ])
console.log(res)

res = makeArrayIncreasing([ 1, 5, 3, 6, 7 ], [ 4, 3, 1 ])
console.log(res)

res = makeArrayIncreasing([ 1, 5, 3, 6, 7 ], [ 1, 6, 3, 3 ])
console.log(res)
 */
res = makeArrayIncreasing([ 0, 11, 6, 1, 4, 3 ], [ 5, 4, 11, 10, 1, 0 ])

console.log(res)
