/**
 * @param {string} S
 * @param {number} K
 * @return {string}
 */
var orderlyQueue = function(S, K) {
  let loop = true
  let res = S
  do {
    loop = false
    index = -1
    let i = 0
    while (res[i] <= res[i + 1] && i < K - 1 && i < res.length - 1) {
      i++
    }

    //console.log(i, res)

    if (res[i] > res[i + 1]) {
      let ch = res[i]
      res = res.replace(ch, '') + ch
      loop = true
    }
  } while (loop)

  return res
}

console.log(orderlyQueue('baaca', 3))
