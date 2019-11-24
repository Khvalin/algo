/**
 * @param {string[]} products
 * @param {string} searchWord
 * @return {string[][]}
 */
var suggestedProducts = function(products, searchWord) {
  products = products.sort((s1, s2) => s1.localeCompare(s2))
  const dict = {}
  for (let i = 0; i < products.length; i++) {
    for (let j = 1; j <= products[i].length; j++) {
      const prefix = products[i].substr(0, j)
      dict[prefix] = dict[prefix] || []

      if (dict[prefix].length < 3) {
        dict[prefix].push(i)
      }
    }
  }

  const res = []
  for (let i = 1; i <= searchWord.length; i++) {
    prefix = searchWord.substr(0, i)
    dict[prefix] = dict[prefix] || []

    res.push(dict[prefix].map((n) => products[n]))
  }

  return res
}

let products = [ 'mobile', 'mouse', 'moneypot', 'monitor', 'mousepad' ],
  searchWord = 'mouse'
console.log(suggestedProducts(products, searchWord))
