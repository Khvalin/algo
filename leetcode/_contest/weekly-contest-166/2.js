/**
 * @param {number[]} groupSizes
 * @return {number[][]}
 */
var groupThePeople = function(groupSizes) {
  let m = {};
  const r = [];

  for (let i = 0; i < groupSizes.length; i++) {
    const g = groupSizes[i];
    m[g] = m[g] || [];

    m[g].push(i);
    if (m[g].length === g) {
      r.push([ ...m[g] ]);
      m[g] = [];
    }
  }

  return r;
};

let groupSizes = [ 3, 3, 3, 3, 3, 1, 3 ];
console.log(groupThePeople(groupSizes));
