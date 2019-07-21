/**
 * @param {number[]} A
 * @return {number}
 */
var numSquarefulPerms = function(A) {
  let squares = {}
  for (let i = 0; i < A.length; i++) {
    for (let j = 0; j < A.length; j++) {
      const s = Math.sqrt(A[i] + A[j])
      if (i != j && Math.floor(s) === Math.ceil(s)) {
        squares[A[i] + A[j]] = true
      }
    }
  }

  const used = {}
  let res = 0
  const check = function(a) {
    const key = a.join(' ')
    if (used[key]) {
      return
    }
    used[key] = true

    for (let i = 1; i < a.length; i++) {
      if (!squares[a[i] + a[i - 1]]) {
        return
      }
    }

    res++
  }

  function swap(alphabets, index1, index2) {
    var temp = alphabets[index1]
    alphabets[index1] = alphabets[index2]
    alphabets[index2] = temp
    return alphabets
  }

  function permute(alphabets, startIndex, endIndex) {
    if (startIndex > 1) {
      if (!squares[alphabets[startIndex] + alphabets[startIndex - 1]]) {
        return
      }
    }

    if (startIndex === endIndex) {
      check(alphabets)
    } else {
      var i
      for (i = startIndex; i <= endIndex; i++) {
        if (startIndex === i || alphabets[startIndex] !== alphabets[i]) {
          swap(alphabets, startIndex, i)
          permute(alphabets, startIndex + 1, endIndex)
          swap(alphabets, i, startIndex) // backtrack
        }
      }
    }
  }

  if (check(A)) {
    res++
  }

  var alphabets = A
  permute(alphabets, 0, alphabets.length - 1) // ABC, ACB, BAC, BCA, CBA, CAB

  //console.log(squares)
  return res
}

let r = numSquarefulPerms([ 1, 17, 8 ])
console.log(r)

r = numSquarefulPerms([ 2, 2, 2 ])
console.log(r)

r = numSquarefulPerms([ 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2 ])
console.log(r)
