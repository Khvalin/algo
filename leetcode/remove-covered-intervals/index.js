/**
 * @param {number[][]} intervals
 * @return {number}
 */
var removeCoveredIntervals = function(intervals) {
  intervals = intervals.sort((a, b) => {
    if (a[0] === b[0]) {
      return b[1] - a[1];
    }
    return a[0] - b[0];
  });

  //  console.log(intervals)

  let j = 0;
  let r = intervals.length > 0 ? 1 : 0;
  for (let i = 1; i < intervals.length; i++) {
    if (intervals[i][0] >= intervals[j][0] && intervals[i][1] <= intervals[j][1]) {
      continue;
    }
    j = i;
    r++;
  }

  return r;
};
