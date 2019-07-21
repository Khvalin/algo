/**
 * @param {number[]} A
 * @return {string}
 */
var largestTimeFromDigits = function(A) {
  function nextPermutation(array) {
    // Find non-increasing suffix
    var i = array.length - 1
    while (i > 0 && array[i - 1] >= array[i]) i--
    if (i <= 0) return false

    // Find successor to pivot
    var j = array.length - 1
    while (array[j] <= array[i - 1]) j--
    var temp = array[i - 1]
    array[i - 1] = array[j]
    array[j] = temp

    // Reverse suffix
    j = array.length - 1
    while (i < j) {
      temp = array[i]
      array[i] = array[j]
      array[j] = temp
      i++
      j--
    }
    return true
  }

  const times = []
  const addTime = (a) => {
    const h = a[0] * 10 + a[1]
    const m = a[2] * 10 + a[3]

    //console.log(times)
    if (h < 24 && m < 60) {
      const t = a[0] + '' + a[1] + ':' + a[2] + '' + a[3]
      times.push(t)
    }
  }

  A.sort()
  do {
    //console.log(A)
    addTime(A)
  } while (nextPermutation(A))

  times.sort()
  //  console.log(times)

  return times.pop() || ''
}

//console.log(largestTimeFromDigits([ 1, 2, 3, 4 ]))

console.log(largestTimeFromDigits([ 1, 9, 6, 0 ]))
