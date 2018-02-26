const getTotalAmoutOfTime = function(timeStapms) {
  const convertToMinutes = (ts) => {
    var int = 1 * ts;
    return Math.trunc(int / 100) * 60 + int % 100;
  };

  const periods = timeStapms
    .map((ts) => {
      return {
        start: convertToMinutes(ts[0]),
        end: convertToMinutes(ts[1])
      };
    })
    .sort((a, b) => (a.start - b.start) * 100 + (a.end - b.end));

  let i = 0;
  while (i < periods.length - 1) {
    const period = periods[i];

    if (periods[i + 1].start < period.end) {
      periods[i + 1].start = Math.min(periods[i + 1].start, period.start);
      periods[i + 1].end = Math.max(periods[i + 1].end, period.end);
      periods.splice(i, 1);
    } else {
      i++;
    }
  }

  const result = periods.reduce((acc, cur) => {
    acc += cur.end - cur.start;
    return acc;
  }, 0);

  return `${Math.trunc(result / 60)}h:${Math.trunc(result % 60)}m`;
};

const result = getTotalAmoutOfTime([ [ '0730', '1200' ], [ '0930', '1330' ], [ '1600', '2000' ], [ '1800', '2400' ] ]);
console.log(result);
