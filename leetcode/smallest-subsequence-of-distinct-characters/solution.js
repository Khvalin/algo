/**
 * @param {string} text
 * @return {string}
 */
var smallestSubsequence = function(text) {
  const abc = []
  let D = {}

  for (const ch of text) {
    D[ch] = D[ch] || 0
    if (!D[ch]) {
      abc.push(ch)
    }
    D[ch]++
  }

  let r = ''
  let dict = {}
  let t = ''
  let c = 0
  for (const ch of text) {
    dict[ch] = dict[ch] || 0
    dict[ch]++
    D[ch]--

    if (dict[ch] > 1) {
      dict[ch]--
      const str = t.replace(ch, '') + ch
      //console.log(t, str)
      if (str.localeCompare(t) < 0) {
        t = str
      }
    } else {
      // if (D[ch] == 1) {
      t += ch
      c++
      //}
    }

    if (c >= abc.length && (t.localeCompare(r) < 0 || r == '')) {
      r = t
    }
  }

  return r
}

console.log(smallestSubsequence('cdadabcc'), 'adbc')
console.log(smallestSubsequence('ecbacba'), 'eacb')

console.log(smallestSubsequence('ddeeeccdce'), 'cde')
