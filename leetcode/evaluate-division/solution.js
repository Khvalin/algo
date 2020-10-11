/**
 * @param {string[][]} equations
 * @param {number[]} values
 * @param {string[][]} queries
 * @return {number[]}
 */
var calcEquation = function(equations, values, queries) {
  const eqs = {};
  const varMap = {};
  for (let i = 0; i < values.length; i++) {
    let [ a, b ] = equations[i];
    eqs[a] = eqs[a] || {};
    eqs[a][b] = values[i];

    eqs[b] = eqs[b] || {};
    eqs[b][a] = 1 / values[i];

    varMap[equations[i][0]] = true;
    varMap[equations[i][1]] = true;
  }

  const vars = Object.keys(varMap);
  for (let i = 0; i < vars.length; i++) {
    let a = vars[i];
    eqs[a] = eqs[a] || {};
    eqs[a][a] = 1;
  }

  let flag = false;
  do {
    flag = false;
    for (let i = 0; i < equations.length; i++) {
      for (let j = i + 1; j < equations.length; j++) {
        let [ a, b ] = equations[i];
        let [ c, d ] = equations[j];

        if (b === c && !eqs[a][d]) {
          eqs[a][d] = eqs[a][b] * eqs[c][d];
          equations.push([ a, d ]);
          flag = true;
        }

        if (a === d && !eqs[b][c]) {
          eqs[b][c] = eqs[b][a] * eqs[d][c];
          equations.push([ b, c ]);
          flag = true;
        }

        [ b, a ] = equations[i];
      }
    }
  } while (flag);
  console.log(equations);

  const res = [];
  for (const q of queries) {
    const [ a, b ] = q;
    if (!varMap[a] || !varMap[b]) {
      res.push(-1);
      continue;
    }

    if (eqs[a][b]) {
      res.push(eqs[a][b]);
      continue;
    }

    res.push(-1);
  }

  return res;
};

let equations = [ [ 'a', 'b' ], [ 'b', 'c' ] ];
let values = [ 2.0, 3.0 ];
let queries = [ [ 'a', 'c' ], [ 'b', 'a' ], [ 'a', 'e' ], [ 'a', 'a' ], [ 'x', 'x' ] ];

equations = [ [ 'a', 'c' ], [ 'b', 'e' ], [ 'c', 'd' ], [ 'e', 'd' ] ];
values = [ 2.0, 3.0, 0.5, 5.0 ];
queries = [ [ 'a', 'b' ] ];

console.log(calcEquation(equations, values, queries));
