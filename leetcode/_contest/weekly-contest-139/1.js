/**
 * @param {string} str1
 * @param {string} str2
 * @return {string}
 */
var gcdOfStrings = function(str1, str2) {
  let res = ''

  for (let i = 0; i <= str1.length; i++) {
    let s = str1.substring(0, i)

    let re = new RegExp(s, 'g')

    if (!str1.replace(re, '').length && !str2.replace(re, '').length) {
      res = s
    }
  }

  return res
}

console.log(gcdOfStrings('ABCABC', 'ABC'))
console.log(gcdOfStrings('ABABAB', 'ABAB'))
