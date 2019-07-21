/**
 * @param {number[]} A
 * @param {number[]} B
 * @return {number[]}
 */
var advantageCount = function(A, B) {
  const bSorted = [...B].sort((a, b) => a - b);
  
  const bIndices = B.map((n, i) => {
    const r = bSorted.indexOf(n);
    bSorted[r] = null;
    return r;
  })

  //console.log (bIndices)
  var aSorted = [...A].sort((a, b) => a - b);
  console.log(aSorted)

  for (let i = 0; i < A.length; i++){
    const pos = bIndices.indexOf(i);
    console.log(pos, B[pos])

    let j = 0;
    while (j < aSorted.length && aSorted[j] <= B[pos]){
      j++;
    }

    j = Math.min(j, aSorted.length - 1)

    console.log(aSorted[j])

    A[pos] = aSorted[j];
    aSorted.splice(j, 1);
  }

  return A
};

console.log(advantageCount([2,7,11,15], [1,10,4,11]))
console.log(advantageCount([12,24,8,32],  [13,25,32,11]))