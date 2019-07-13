//Generate a random number between 1-7 given a function that returns a random number between 1-5
function rand5() {
  return Math.floor(Math.random() * 5) + 1
}

function rand7() {
  let sum = 0

  do {
    sum += rand5()
  } while (sum % 8 > 0)

  return sum % 7 + 1
}

const hash = {}
for (let i = 0; i < 1000; i++) {
  const rnd = rand7()
  hash[rnd] = hash[rnd] || 0
  hash[rnd]++
}

console.log(hash)
