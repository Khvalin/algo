/**
 * @param {string[]} words
 */
var StreamChecker = function(words) {
  this.words = words
  //this.indices = words.map((w) => w.length - 1)
  this.count = words.map((w) => 0)
}

/** 
 * @param {character} letter
 * @return {boolean}
 */
StreamChecker.prototype.query = function(letter) {
  let r = false
  for (let i = 0; i < this.copy.length; i++) {
    if (this.copy[i][this.count[i]] == letter) {
      this.copy[i] = this.copy[i].substr(1)

      if (this.copy[i] == '') {
        r = true
      }
    }
  }

  return r
}

/** 
 * Your StreamChecker object will be instantiated and called as such:
 * var obj = new StreamChecker(words)
 * var param_1 = obj.query(letter)
 */

var obj = new StreamChecker([ 'ab', 'ba', 'aaab', 'abab', 'baa' ]) // init the dictionary.

const QUERIES = [
  [ 'a' ],
  [ 'a' ],
  [ 'a' ],
  [ 'a' ],
  [ 'a' ],
  [ 'b' ],
  [ 'a' ],
  [ 'b' ],
  [ 'a' ],
  [ 'b' ],
  [ 'b' ],
  [ 'b' ],
  [ 'a' ],
  [ 'b' ],
  [ 'a' ],
  [ 'b' ],
  [ 'b' ],
  [ 'b' ],
  [ 'b' ],
  [ 'a' ],
  [ 'b' ],
  [ 'a' ],
  [ 'b' ],
  [ 'a' ],
  [ 'a' ],
  [ 'a' ],
  [ 'b' ],
  [ 'a' ],
  [ 'a' ]
  // [ 'a' ]
].map((a) => obj.query(a[0]))

console.log(QUERIES)
//console.log(obj.copy)
