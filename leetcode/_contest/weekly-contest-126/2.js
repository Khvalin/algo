/**
 * @param {string} S
 * @return {boolean}
 */
var isValid = function(S) {
  const ABC = 'abc'
  while (S.indexOf(ABC) > -1) {
    S = S.replace(ABC, '')
  }

  return S.length === 0
}

console.log(isValid('aabcbc'))
console.log(isValid('abcabcababcc'))
console.log(isValid('abccba'))
