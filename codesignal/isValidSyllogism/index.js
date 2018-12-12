const isValidSyllogism = (syllogism) => {
  const COUNT = 4
  const sets = {}

  const next = (function(n) {
    return () => ++n
  })(0)

  const SOME = (a, b) => {
    let [ aVal, bVal ] = [ sets[a] || [], sets[b] || [] ]
    let i = 0
    while (aVal.length < COUNT / 2 || bVal.length < COUNT / 2) {
      let n = aVal[i] || bVal[i] || next() //TODO
      if (!aVal[i]) {
        aVal.push(n)
      }
      if (!bVal[i]) {
        bVal.push(n)
      }
      i++
    }

    while (aVal.length < COUNT || bVal.length < COUNT) {
      if (!aVal[i]) {
        aVal.push(next())
      }
      if (!bVal[i]) {
        bVal.push(next())
      }
      i++
    }

    sets[a] = aVal
    sets[b] = bVal
  }

  const clauses = [
    {
      regex: /All ([A-Z]) is ([A-Z])/,
      checkConclusion: (a, b) => {
        let [ aVal, bVal ] = [ sets[a] || [], sets[b] || [] ]
        return aVal.every((cur) => bVal.indexOf(cur) > -1)
      },
      parse: (a, b) => {
        let [ aVal, bVal ] = [ sets[a] || [], sets[b] || [] ]
        let i = 0
        while (aVal.length < COUNT || bVal.length < COUNT) {
          let n = aVal[i] || bVal[i] || next() //TODO
          if (!aVal[i]) {
            aVal.push(n)
          }
          if (!bVal[i]) {
            bVal.push(n)
          }
          i++
        }

        sets[a] = aVal
        sets[b] = bVal
      }
    },
    {
      regex: /No ([A-Z]) is ([A-Z])/,
      checkConclusion: (a, b) => {
        let [ aVal, bVal ] = [ sets[a] || [], sets[b] || [] ]
        return !aVal.some((cur) => bVal.indexOf(cur) > -1)
      },
      parse: (a, b) => {
        let [ aVal, bVal ] = [ sets[a] || [], sets[b] || [] ]

        while (aVal.length < COUNT || bVal.length < COUNT) {
          let n = next()
          if (aVal.length < bVal.length) {
            if (bVal.indexOf(n) === -1) {
              aVal.push(n)
            }
          } else {
            if (aVal.indexOf(n) === -1) {
              bVal.push(n)
            }
          }
        }

        sets[a] = aVal
        sets[b] = bVal
      }
    },
    {
      regex: /Some ([A-Z]) is ([A-Z])/,
      checkConclusion: (a, b) => {
        let [ aVal, bVal ] = [ sets[a] || [], sets[b] || [] ]
        return aVal.some((cur) => bVal.indexOf(cur) > -1)
      },
      parse: SOME
    },
    {
      regex: /Some ([A-Z]) is not ([A-Z])/,
      checkConclusion: (a, b) => {
        let [ aVal, bVal ] = [ sets[a] || [], sets[b] || [] ]
        return !aVal.every((cur) => bVal.indexOf(cur) > -1)
      },
      parse: (a, b) => SOME
    }
  ]

  let result = false
  for (let i = 0; i < syllogism.length; i++) {
    for (const clause of clauses) {
      const matches = syllogism[i].match(clause.regex)
      if (matches) {
        if (i < 2) {
          clause.parse(matches[1], matches[2])
        } else {
          result = clause.checkConclusion(matches[1], matches[2])
        }
        //  console.log([ ...matches ])
      }
    }
  }

  console.log(sets)
  return result
}

console.log(isValidSyllogism([ 'No W is X', 'All P is W', 'Some P is not X' ]))
console.log(isValidSyllogism([ 'All P is W', 'No P is X', 'Some P is not X' ]))
console.log(isValidSyllogism([ 'Some P is W', 'No P is X', 'Some P is not X' ]))

console.log(isValidSyllogism([ 'All Q is F', 'No O is F', 'No O is Q' ]))
console.log(isValidSyllogism([ 'Some E is T', 'All Q is T', 'Some Q is E' ]))

//console.log('All P is W'.match(/All ([A-Z]) is ([A-Z])/))
