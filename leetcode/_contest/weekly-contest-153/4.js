/**
 * @param {number[]} arr1
 * @param {number[]} arr2
 * @return {number}
 */
var makeArrayIncreasing = function(arr1, arr2) {
  const memo = {}

  const calc = (k, l, prev = Number.NEGATIVE_INFINITY) => {
    let r1 = Number.POSITIVE_INFINITY
    let r2 = Number.POSITIVE_INFINITY

    if (k >= arr1.length) {
      return 0
    }

    const key = `${k}_${l}_${prev}`
    if (key in memo) {
      return memo[key]
    }

    while (l < arr2.length && arr2[l] <= prev) {
      l++
    }

    if (l < arr2.length) {
      r1 = 1 + calc(k + 1, l + 1, arr2[l])
    }

    if (arr1[k] > prev) {
      r2 = calc(k + 1, l, arr1[k])
    }

    const r = Math.min(r1, r2)
    memo[key] = r

    return r
  }

  arr2 = arr2.sort((a, b) => a - b)

  const res = calc(0, 0)

  return res > arr1.length ? -1 : res
}

let res
/* 
res = makeArrayIncreasing([ 1, 5, 3, 6, 7 ], [ 1, 3, 2, 4 ])
console.log(res)



res = makeArrayIncreasing([ 1, 5, 3, 6, 7 ], [ 1, 6, 3, 3 ])
console.log(res)
 */

res = makeArrayIncreasing([ 1, 5, 3, 6, 7 ], [ 1, 6, 3, 3 ])
console.log(res)

res = makeArrayIncreasing([ 1, 2, 3, 4, 4 ], [ 0, 1, 2, 3 ])
console.log(res)

res = makeArrayIncreasing([ 1, 5, 3, 6, 7 ], [ 4, 3, 1 ])
console.log(res)

res = makeArrayIncreasing([ 0, 11, 6, 1, 4, 3 ], [ 5, 4, 11, 10, 1, 0 ])
console.log(res)
