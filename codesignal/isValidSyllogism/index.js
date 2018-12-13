const isValidSyllogism = (syllogism) => {
  const INTERSECTION = {}
  const EMPTY_SET = '∅'
  const NON_EMPTY_SET = '~'
  const NON_INCLUDING = '%'

  const clauses = [
    {
      regex: /All ([A-Z]) is ([A-Z])/,
      checkConclusion: (a, b) => {},
      parse: (a, b) => {
        INTERSECTION[a + b] = a
        INTERSECTION[b + a] = a
      }
    },
    {
      regex: /No ([A-Z]) is ([A-Z])/,
      checkConclusion: (a, b) => {},
      parse: (a, b) => {
        INTERSECTION[a + b] = EMPTY_SET
        INTERSECTION[b + a] = EMPTY_SET
      }
    },
    {
      regex: /Some ([A-Z]) is ([A-Z])/,
      checkConclusion: (a, b) => {},
      parse: (a, b) => {
        INTERSECTION[a + b] = NON_EMPTY_SET
        INTERSECTION[b + a] = NON_EMPTY_SET
      }
    },
    {
      regex: /Some ([A-Z]) is not ([A-Z])/,
      checkConclusion: (a, b) => {},
      parse: (a, b) => {
        INTERSECTION[a + b] = NON_INCLUDING
        //INTERSECTION[b + a] = NON_EMPTY_SET
      }
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

  console.log(INTERSECTION)

  return result
}

console.log(isValidSyllogism([ 'All Q is F', 'No O is F', 'No O is Q' ]))
/*
console.log(isValidSyllogism([ 'No W is X', 'All P is W', 'Some P is not X' ]))
console.log(isValidSyllogism([ 'All P is W', 'No P is X', 'Some P is not X' ]))
console.log(isValidSyllogism([ 'Some P is W', 'No P is X', 'Some P is not X' ]))


console.log(isValidSyllogism([ 'Some E is T', 'All Q is T', 'Some Q is E' ]))
*/
//console.log('All P is W'.match(/All ([A-Z]) is ([A-Z])/))
