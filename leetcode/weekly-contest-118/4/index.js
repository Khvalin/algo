/**
 * @param {string} S
 * @param {string} T
 * @return {boolean}
 */
var isRationalEqual = function(S, T) {
  let parse = (str) => {
    let [ int, floatPart ] = str.split('.')
    int = int || '0'
    floatPart = floatPart || '0'

    let left = floatPart.indexOf('(')
    let right = floatPart.indexOf(')')
    if (left > -1) {
      let period = floatPart.substring(left + 1, right)

      //console.log(period)

      let rest = floatPart.substring(right + 1)
      floatPart = floatPart.substring(0, left)

      floatPart += period.repeat(10)
      floatPart = floatPart.substring(0, 9)

      if (!floatPart.endsWith(rest)) {
        floatPart += rest
      }
    }

    let res = 0
    res += parseInt(int)

    if (floatPart) {
      res += parseFloat('0.' + floatPart)
      //   console.log(floatPart, parseFloat(floatPart))
    }

    return res
  }

  let [ s, t ] = [ parse(S), parse(T) ]

  console.log(s, t, Math.abs(s - t))

  return Math.abs(s - t) <= 1.0e-9
}

//console.log(isRationalEqual('7.(52)', '7.5(25)'))
//console.log(isRationalEqual('1.0', '0.(99)'))
console.log(isRationalEqual('1.0', '0.9(9)'))
//console.log(isRationalEqual('1.0', '0.(99)1'))
