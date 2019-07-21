/**
 * @param {number[]} barcodes
 * @return {number[]}
 */
var rearrangeBarcodes = function(barcodes) {
  const unique = (a) => {
    const m = {}
    let r = []
    for (n of a) {
      if (!m[n]) {
        r.push(n)
        m[n] = 1
      }
    }

    return r
  }

  let count = {}

  for (b of barcodes) {
    count[b] = count[b] || 0
    count[b]++
  }

  let u = unique([ ...barcodes ]).sort((a, b) => count[a] < count[b])

  const r = []
  let i = 0,
    p = -1
  while (r.length < barcodes.length) {
    let n = u[i % u.length]
    i = (i + 1) % 2
    if (n == p) {
      continue
    }

    r.push(n)
    count[n]--
    if (count[n] == 0) {
      u = u.filter((a) => a != n)
    }
    u = u.sort((a, b) => count[a] < count[b])

    p = n
  }

  return r
}

let ans

ans = rearrangeBarcodes([ 1, 1, 1, 2, 2, 2 ])
console.log(ans)

ans = rearrangeBarcodes([ 1, 1, 1, 1, 2, 2, 3, 3 ])
console.log(ans)

// prettier-ignore
ans = rearrangeBarcodes([51,83,51,40,51,40,51,40,83,40,83,83,51,40,40,51,51,51,40,40,40,83,51,51,40,51,51,40,40,51,51,40,51,51,51,40,83,40,40,83,51,51,51,40,40,40,51,51,83,83,40,51,51,40,40,40,51,40,83,40,83,40,83,40,51,51,40,51,51,51,51,40,51,83,51,83,51,51,40,51,40,51,40,51,40,40,51,51,51,40,51,83,51,51,51,40,51,51,40,40])
console.log(ans.join(' '))
