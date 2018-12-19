const isValidSyllogism = (syllogism) => {

  const INTERSECTION = {}
  const EMPTY_SET = '∅'
  const NON_EMPTY_SET = '~'
  const NON_INCLUDING = '%'

  const allIsCheck = (a, b, c) => {
    return (INTERSECTION[a + b] === a)
      || (INTERSECTION[a + c] === a && INTERSECTION[c + b] === b)

  }

  const noIsCheck = (a, b, c) => {
    return INTERSECTION[a + b] === EMPTY_SET
      || (INTERSECTION[a + c] === a && INTERSECTION[c + b] === EMPTY_SET)
      || (INTERSECTION[b + c] === b && INTERSECTION[c + a] === EMPTY_SET)
      || (INTERSECTION[a + c] === a && INTERSECTION[c + b] === EMPTY_SET)

      ;
  }

  const clauses = [
    {
      regex: /All ([A-Z]) is ([A-Z])/,
      checkConclusion: allIsCheck,
      parse: (a, b) => {
        INTERSECTION[a + b] = a
        INTERSECTION[b + a] = a
      }
    },
    {
      regex: /No ([A-Z]) is ([A-Z])/,
      checkConclusion: noIsCheck,
      parse: (a, b) => {
        INTERSECTION[a + b] = EMPTY_SET
        INTERSECTION[b + a] = EMPTY_SET
      }
    },
    {
      regex: /Some ([A-Z]) is ([A-Z])/,
      checkConclusion: (a, b, c) => {
        return allIsCheck(a, b, c)
          || (INTERSECTION[a + c] == NON_EMPTY_SET && INTERSECTION[c + b] == c)
          || (INTERSECTION[c + a] == NON_EMPTY_SET && INTERSECTION[c + b] == c)
          || (INTERSECTION[b + c] == NON_EMPTY_SET && INTERSECTION[c + a] == c)
      },
      parse: (a, b) => {
        INTERSECTION[a + b] = NON_EMPTY_SET
        INTERSECTION[b + a] = NON_EMPTY_SET
      }
    },
    {
      regex: /Some ([A-Z]) is not ([A-Z])/,
      checkConclusion: (a, b, c) => {
        return noIsCheck(a, b, c)
          || (INTERSECTION[a + c] == NON_EMPTY_SET && (INTERSECTION[c + b] == NON_INCLUDING || INTERSECTION[c + a] == EMPTY_SET))
          || (INTERSECTION[c + a] == NON_EMPTY_SET && (INTERSECTION[c + b] == NON_INCLUDING || INTERSECTION[c + b] == EMPTY_SET))
          || (INTERSECTION[b + c] == NON_EMPTY_SET && (INTERSECTION[c + a] == NON_INCLUDING || INTERSECTION[c + a] == EMPTY_SET))
        //|| (INTERSECTION[a + c] == NON_EMPTY_SET && INTERSECTION[c + b]
      },
      parse: (a, b) => {
        INTERSECTION[a + b] = NON_INCLUDING
        //INTERSECTION[b + a] = NON_EMPTY_SET
      }
    }
  ]

  const setNames = new Set()
  let result = false
  for (let i = 0; i < syllogism.length; i++) {
    for (const clause of clauses) {
      const matches = syllogism[i].match(clause.regex)
      if (matches) {
        if (i < 2) {
          setNames.add(matches[2])
          setNames.add(matches[1])
          clause.parse(matches[1], matches[2])
        } else {
          setNames.delete(matches[2])
          setNames.delete(matches[1])
          const third = [...setNames.values()][0]
          result = clause.checkConclusion(matches[1], matches[2], third)
        }
        //  console.log([ ...matches ])
      }
    }
  }

  // console.log(INTERSECTION)

  return result
}

const tests = [
  {
    input: ["No W is X",
      "All P is W",
      "Some P is not X"
    ], result: true
  },
  {
    input: ["All Q is F",
      "No O is F",
      "No O is Q"], result: true
  },

  {
    input: ["Some E is T",
      "All Q is T",
      "Some Q is E"], result: false
  },
  {
    input: ["No A is J",
      "All A is R",
      "All R is J"], result: false
  },
  {
    input: ["All F is K",
      "Some Z is not K",
      "Some Z is not F"], result: true
  },

];

for (let i = 0; i < tests.length; i++) {
  const t = tests[i]
  const res = isValidSyllogism(t.input)
  let message = `Test ${i} passed`
  if (res !== t.result) {
    message = `Test ${i} failed. For ${t.input} expected ${t.result}, got ${res}`
  }
  console.log(message)
}