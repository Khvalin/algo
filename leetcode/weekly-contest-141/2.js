/**
 * @param {number[]} values
 * @param {number[]} labels
 * @param {number} num_wanted
 * @param {number} use_limit
 * @return {number}
 */
var largestValsFromLabels = function(values, labels, num_wanted, use_limit) {
  let items = values.map((n, i) => {
    return { num: n, label: labels[i] }
  })

  items = items.sort((a, b) => b.num - a.num)

  let [ count, res ] = [ 0, 0 ]
  let labs = {}

  for (let i = 0; count < num_wanted && i < items.length; i++) {
    labs[items[i].label] = labs[items[i].label] || 0
    labs[items[i].label]++
    //
    if (labs[items[i].label] <= use_limit) {
      count++
      res += items[i].num
    }
  }

  return res
}

let res = largestValsFromLabels([ 5, 4, 3, 2, 1 ], [ 1, 1, 2, 2, 3 ], 3, 1)
console.log(res)

res = largestValsFromLabels([ 9, 8, 8, 7, 6 ], [ 0, 0, 0, 1, 1 ], 3, 2)
console.log(res)
