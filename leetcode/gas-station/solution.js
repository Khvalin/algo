/**
 * @param {number[]} gas
 * @param {number[]} cost
 * @return {number}
 */
/**
 * @param {number[]} gas
 * @param {number[]} cost
 * @return {number}
 */
var canCompleteCircuit = function(gas, cost) {
  let start = 0;
  let g = Number.MIN_VALUE;

  for (let i = 0; i < gas.length; i++) {
    const d = gas[i] - cost[i];
    if (d > g) {
      start = i;
      g = d;
    }
  }

  if (g < 0) {
    return -1;
  }

  // start = 3;

  let j = start;
  console.log(j);
  for (let i = (start + 1) % gas.length; i !== start; i = (i + 1) % gas.length) {
    if (g < 0) {
      return -1;
    }

    g -= cost[j];

    if (g < 0) {
      return -1;
    }

    g += gas[i];
    j = i;
  }

  return start;
};

var canCompleteCircuit = function(gas, cost) {
  let cur = 0;
  let g = 0;
  let j = 0;
  for (let i = 0; i < gas.length; i++) {
    cur += gas[i] - cost[i];
    g += gas[i] - cost[i];
    if (cur < 0) {
      cur = 0;
      j = i + 1;
    }
  }
  return g < 0 ? -1 : j;
};

let gas = [ 1, 2, 3, 4, 5 ];
let cost = [ 3, 4, 5, 1, 2 ];
let r = null;

gas = [ 2, 3, 4 ];
cost = [ 3, 4, 3 ];

gas = [ 5, 8, 2, 8 ];
cost = [ 6, 5, 6, 6 ];

r = canCompleteCircuit(gas, cost);

console.log(r);
