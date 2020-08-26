/**
 * @param {number[]} days
 * @param {number[]} costs
 * @return {number}
 */
const mincostTickets = (days, costs) => {
  let res = Number.POSITIVE_INFINITY;

  const a = [].fill(Number.POSITIVE_INFINITY, 0, 500);
  const q = [ { day: 0, total: 0 } ];
  const last = days[days.length - 1];

  while (q.length > 0) {
    let { day, total } = q.shift();

    if (day > last) {
      res = Math.min(res, total);
      continue;
    }

    // console.log(day, total);

    if (total >= a[day]) {
      continue;
    }

    a[day] = total;

    let i = 0;
    while (i < days.length - 1 && day > days[i]) {
      i++;
    }

    day = days[i];
    q.push(
      { day: day + 1, total: total + costs[0] },
      { day: day + 7, total: total + costs[1] },
      { day: day + 30, total: total + costs[2] }
    );
  }

  return res;
};

console.log(mincostTickets([ 1, 4, 6, 7, 8, 20 ], [ 2, 7, 15 ]));
console.log(mincostTickets([ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31 ], [ 2, 7, 15 ]));
console.log(mincostTickets([ 1, 2, 3, 4 ], [ 2, 7, 15 ]));
