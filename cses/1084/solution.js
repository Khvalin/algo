var readline = require("readline");

var r = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
  terminal: false,
});

const readdata = new Promise((resolve) => {
  const covert = (s) => {
    return s.split(" ").map((n) => 1 * n);
  };

  const lines = [];

  r.on("line", function (line) {
    lines.push(covert(line));
  });
  r.on("close", () => {
    resolve(lines);
  });
});

const solve = (lines) => {
  const [_2, _1, k] = lines[0];
  const apartments = lines[1];
  const applicants = lines[2];

  apartments.sort((a, b) => a - b);
  applicants.sort((a, b) => a - b);

  let i = apartments.length - 1;
  let res = 0;

  for (let j = applicants.length - 1; j >= 0; j--) {
    const a = applicants[j];
    while (i >= 0 && apartments[i] > a + k) {
      i--;
    }
    if (i < 0) {
      break;
    }

    if (apartments[i] <= a + k && apartments[i] >= a - k) {
      i--;
      res++;
    }
  }

  return res;
};

const writedata = (res) => {
  console.log(res);
};

readdata.then(solve).then(writedata);
