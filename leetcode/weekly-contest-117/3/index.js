/**
 * @param {string[]} wordlist
 * @param {string[]} queries
 * @return {string[]}
 */
var spellchecker = function(wordlist, queries) {
  const replaceVowels = (w = '') => {
    //  const vowelsRegEx = /[aeiou]/gi
    return w.toUpperCase().replace(/[aeiou]/gi, '~')
  }

  const dict = {}
  const caseCheck = {}
  const vowelCheck = {}

  for (let i = 0; i < wordlist.length; i++) {
    let w = wordlist[i]

    if (!(w in dict)) {
      dict[w] = i
    }

    w = w.toUpperCase()

    if (!(w in caseCheck)) {
      caseCheck[w] = i
    }

    const noVowels = replaceVowels(w)
    if (!(noVowels in vowelCheck)) {
      vowelCheck[noVowels] = i
    }
  }

  //console.log(vowelCheck)

  const result = queries.map((q = '') => {
    if (q in dict) {
      return wordlist[dict[q]]
    }

    q = q.toUpperCase()

    if (q in caseCheck) {
      return wordlist[caseCheck[q]]
    }

    q = replaceVowels(q)

    if (q in vowelCheck) {
      return wordlist[vowelCheck[q]]
    }

    return ''
  })

  return result
}

let wordlist = [ 'KiTe', 'kite', 'hare', 'Hare' ]
let queries = [ 'kite', 'Kite', 'KiTe', 'Hare', 'HARE', 'Hear', 'hear', 'keti', 'keet', 'keto' ]

console.log(spellchecker(wordlist, queries))
