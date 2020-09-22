/**
 * @param {number[][]} trips
 * @param {number} capacity
 * @return {boolean}
 */
const carPooling = (trips, capacity) => {
  trips = trips.sort((a, b) => {
    return a[1] - b[1];
  });

  const unload = [];
  for (let i = 0; i < 1001; i++) {
    unload.push(0);
  }

  let cur = 0;
  let prev = 0;
  for (const trip of trips) {
    for (let i = prev; i <= trip[1]; i++) {
      cur -= unload[i];
    }
    cur += trip[0];
    if (cur > capacity) {
      return false;
    }

    unload[trip[2]] += trip[0];
    prev = trip[1] + 1;
  }

  return true;
};
