/**
 * @param {number} N
 * @return {number}
 */
var primePalindrome = function(N) {
  function nextPalindrome(n) {
    let digits = []
    let m = n
    while (m > 0) {
      digits.unshift(m % 10)
      m = Math.trunc(m / 10)
    }

    let carryFlag = false
    let midPoint = Math.floor(digits.length / 2)

    for (let i = 0; i < digits.length; i++) {
      let copyFrom = digits[i]
      let copyTo = digits[digits.length - 1 - i] + carryFlag
      if (copyTo === 10) {
        digits[digits.length - 1 - i] = 0
        carryFlag = true
      } else {
        let digit = i < midPoint ? copyFrom : copyTo
        carryFlag = copyTo > digit
        digits[i] = digits[digits.length - 1 - i] = digit
      }
    }

    return digits.reduce((acc, next) => 10 * acc + next, 0)
  }

  const primes = {}

  const isPrime = (num) => {
    if (num in primes) {
      return primes[num]
    }

    for (let i = 2, s = Math.sqrt(num); i <= s; i++) {
      if (num % i === 0) {
        primes[num] = false
        return false
      }
    }

    primes[num] = num !== 1
    return num !== 1
  }

  let next = nextPalindrome(N)
  while (!isPrime(next)) {
    let next = nextPalindrome(N)
    N = next + 1
  }

  return next
}

console.log(primePalindrome(10))
