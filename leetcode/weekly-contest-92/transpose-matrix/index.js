/**
 * @param {number[][]} A
 * @return {number[][]}
 */
var transpose = function(A) {
  var result = []

  for (var i = 0; i < A[0].length; i++) {
    result.push(new Array(A.length))
  }

  for (var i = 0; i < A.length; i++) {
    for (var j = 0; j < A[i].length; j++) {
      result[j][i] = A[i][j]
    }
  }

  return result
}

console.log(transpose([ [ 1, 2, 3 ], [ 4, 5, 6 ], [ 7, 8, 9 ] ]))
console.log(transpose([ [ 1, 2, 3 ], [ 4, 5, 6 ] ]))
