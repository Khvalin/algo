/**
 * @param {number[]} status
 * @param {number[]} candies
 * @param {number[][]} keys
 * @param {number[][]} containedBoxes
 * @param {number[]} initialBoxes
 * @return {number}
 */
var maxCandies = function(status, candies, keys, containedBoxes, initialBoxes) {
  const solve = (boxes = [], ks = []) => {
    if (boxes.length === 0) {
      return 0;
    }

    let nextBoxes = [];
    let nextKeys = [ ...ks ];

    let goNext = false;
    for (const i of boxes) {
      for (const j of keys[i]) {
        if (!nextKeys[j]) {
          goNext = true;
        }
        nextKeys[j] = true;
      }
    }

    let s = 0;
    for (const i of boxes) {
      let st = status[i] === 1 || ks[i];

      status[i] = st ? 1 : 0;

      if (st) {
        s += candies[i];
        if (containedBoxes[i].length) {
          goNext = true;
        }
        nextBoxes = [ ...nextBoxes, ...containedBoxes[i] ];
        containedBoxes[i] = [];
      } else {
        nextBoxes.push(i);
      }
    }

    if (goNext) {
      s += solve(nextBoxes, nextKeys);
    }

    return s;
  };

  return solve(initialBoxes, initialBoxes.map(() => false));
};
