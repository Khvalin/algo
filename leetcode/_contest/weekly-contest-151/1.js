/**
 * @param {string[]} transactions
 * @return {string[]}
 */
var invalidTransactions = function(transactions) {
  const cities = {}
  let t = 0
  const res = []
  for (const tr of transactions) {
    const [ name, time, amount, city ] = tr.split(',')
    t += time

    if (amount > 1000 || (cities[city] > 0 && t - cities[city] <= 60)) {
      res.push(tr)
    }
    cities[city] = t
  }

  return res
}

let transactions = [ 'alice,20,800,mtv', 'alice,50,100,beijing' ]
console.log(invalidTransactions(transactions))

transactions = [ 'alice,20,800,mtv', 'alice,50,1200,mtv' ]
console.log(invalidTransactions(transactions))
