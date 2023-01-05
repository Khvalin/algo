var readline = require("readline");

var r = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
  terminal: false,
});

const readdata = new Promise((resolve) => {
  const convert = (s) => {
    return s.split(/\s+/).map((n) => 1 * n);
  };

  const lines = [];

  r.on("line", function (line) {
    const nums = convert(line) || [];
    if (nums.length) {
      lines.push(nums);
    }
  });
  r.on("close", () => {
    resolve(lines);
  });
});

const solve = (lines) => {
  const [_, x] = lines[0];
  const weights = lines[1];

  weights.sort((a, b) => a - b);
  let i = 0;
  let j = weights.length - 1;

  let res = weights.length;
  while (i < j) {
    if (weights[i] + weights[j] <= x) {
      i++;
      j--;
      res--;
      continue;
    }

    j--;
  }

  return res;
};

const writedata = (res) => {
  console.log(res);
};

readdata.then(solve).then(writedata);
