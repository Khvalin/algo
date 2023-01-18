var ProductOfNumbers = function () {
  const a = (new Array(101)).fill(0);
  this.total = [...a];
  this.counts = [[...a]];
};

/**
 * @param {number} num
 * @return {void}
 */
ProductOfNumbers.prototype.add = function (num) {
  this.total[num]++;
  this.counts.push([...this.total]);
};

/**
 * @param {number} k
 * @return {number}
 */
ProductOfNumbers.prototype.getProduct = function (k) {
  let res = 1;
  const j = this.counts.length - k - 1;
  for (let i = 0; i <= 100; i++) {
    const c = this.total[i] - this.counts[j][i];
    if (c === 0) {
      continue;
    }
    if (i === 0) {
      return 0;
    }
    res *= i ** c;
  }

  return res;
};

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * var obj = new ProductOfNumbers()
 * obj.add(num)
 * var param_2 = obj.getProduct(k)
 */

let methods = [
  "ProductOfNumbers",
  "add",
  "getProduct",
  "getProduct",
  "add",
  "add",
  "getProduct",
  "add",
  "getProduct",
  "add",
  "getProduct",
  "add",
  "getProduct",
  "getProduct",
  "add",
  "getProduct",
];
let params = [
  [],
  [7],
  [1],
  [1],
  [4],
  [5],
  [3],
  [4],
  [4],
  [3],
  [4],
  [8],
  [1],
  [6],
  [2],
  [3],
];

var obj = new ProductOfNumbers();
const a = [];
for (let i = 1; i < params.length; i++) {
  const res = obj[methods[i]](...params[i]);
  if (res === undefined) {
    a.push(...params[i]);
    continue;
  }
  console.log(a);
  console.log(methods[i], params[i], res);
}
