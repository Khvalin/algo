/**
 * @param {string} characters
 * @param {number} combinationLength
 */
var CombinationIterator = function(characters, combinationLength) {
  this.chars = characters;
  this.indices = [];
  this.cLength = combinationLength;

  for (let i = 0; i < combinationLength; i++) {
    this.indices.push(i);
  }
};

/**
 * @return {string}
 */
CombinationIterator.prototype.next = function() {
  let res = '';
  for (const i of this.indices) {
    res += this.chars[i];
  }

  let i = this.indices.length - 1;
  let skip = true;
  for (; i >= 0 && skip; i--) {
    const ind = this.indices[i];
    skip = ind === this.chars.length - 1 || this.indices[i + 1] === ind + 1;
  }

  this.indices[i]++;
  for (let j = i + 1; j < this.indices.length; j++) {
    this.indices[j] = this.indices[j - 1] + 1;
  }

  return res;
};

/**
 * @return {boolean}
 */
CombinationIterator.prototype.hasNext = function() {
  for (const i of this.indices) {
    if (i <= this.chars.length - this.cLength) {
      return true;
    }
  }

  return false;
};

var iter;
iter = new CombinationIterator('abcd', 2);
while (iter.hasNext()) {
  console.log(iter.next());
}

/** 
 * Your CombinationIterator object will be instantiated and called as such:
 * var obj = new CombinationIterator(characters, combinationLength)
 * var param_1 = obj.next()
 * var param_2 = obj.hasNext()
 */
