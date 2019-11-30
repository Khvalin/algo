/**
 * @param {string} num
 * @return {string}
 */
var toHexspeak = function(num) {
  const s = parseInt(num).toString(16).toUpperCase()
  let res = ''

  for (const ch of s) {
    if (ch === '1') {
      res += 'I'
      continue
    }
    if (ch === '0') {
      res += 'O'
      continue
    }

    if ('ABCDEFIO'.indexOf(ch) < 0) {
      res = 'ERROR'
      break
    }
    res += ch
  }
  return res
}

//console.log(toHexspeak(257))
//console.log(toHexspeak(3))
console.log(toHexspeak('747823223228'))
