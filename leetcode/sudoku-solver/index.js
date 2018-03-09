const arrayDiff = function(source, ...arrays) {
  return source.filter(function(i) {
    for (let j = 0; j < arrays.length; j++) {
      if (arrays[j].indexOf(i) > -1) {
        return false;
      }
    }
    return true;
  });
};

class Sudoku {
  constructor(board) {
    this.board = board;
    this.indices = board.map((line, i) => i);
    this.digits = this.indices.map((i) => (i + 1).toString());
  }

  getColumn(index) {
    return this.board.map((line) => line[index]);
  }

  getColumnCandidates(index) {
    return this.getColumn(index);
  }

  getQuadrantCandidates(row, column) {
    const result = [];
    for (let i = 0; i < 3; i++) {
      for (let j = 0; j < 3; j++) {
        if (this.board[3 * Math.floor(row / 3) + i][3 * Math.floor(column / 3) + j] !== '.') {
          result.push(this.board[3 * Math.floor(row / 3) + i][3 * Math.floor(column / 3) + j]);
        }
      }
    }

    return result;
  }

  generateCandidates() {
    const columnCandidates = this.indices.map((i) => this.getColumnCandidates(i));

    return this.board.map((row, i) =>
      row.map((cell, j) => arrayDiff(this.digits, this.getQuadrantCandidates(i, j), columnCandidates[j], row))
    );
  }

  get candidates() {
    this._candidates = this._candidates || this.generateCandidates();
    return this._candidates;
  }

  resetCaches() {
    this._candidates = null;
  }

  isSquareEmpty(rIndex, cIndex) {
    return this.board[rIndex][cIndex] === '.';
  }

  setValue(rIndex, cIndex, val) {
    this.board[rIndex][cIndex] = val;
    this.resetCaches();
    /*
    this.indices.map((i) => {
      this._candidates[i][cIndex] = this._candidates[i][cIndex].filter((v) => val !== v);
      this._candidates[rIndex][i] = this._candidates[rIndex][i].filter((v) => val !== v);
    });
    */
  }

  clearValue(rIndex, cIndex) {
    const val = this.board[rIndex][cIndex];
    this.board[rIndex][cIndex] = '.';
    this.resetCaches();
    /*
    this.indices.map((i) => {
      this._candidates[i][cIndex].push(val);
      this._candidates[rIndex][i].push(val);
    });*/
  }

  solve() {
    const emptySquares = this.indices.reduce((acc, cur, i) => {
      this.indices.map((j) => {
        if (this.isSquareEmpty(i, j)) {
          acc.push({
            rIndex: i,
            cIndex: j
            // priority: 1
            // tried: [],
            // candidates: this.candidates[i][j],
            // weight: this.candidates[i][j].length
          });
        }
      });

      return acc;
    }, []);

    const getSmallest = () => {
      return emptySquares.reduce((lowest, next) => {
        return this.isSquareEmpty(next.rIndex, next.cIndex) &&
        (!lowest ||
          this.candidates[next.rIndex][next.cIndex].length <
            /*          */ this.candidates[lowest.rIndex][lowest.cIndex].length)
          ? next
          : lowest;
      }, null);
    };

    const history = [];
    let count = 0;
    let rollBack = { square: null, attempts: [] };

    //  console.log(this.candidates);

    //  process.exit();

    while (emptySquares.length > count) {
      let attempts = rollBack.attempts;
      const cur = rollBack.square || getSmallest();

      const c = cur && arrayDiff(this.candidates[cur.rIndex][cur.cIndex], attempts);
      const val = c && c[0]; // c[Math.floor(Math.random() * c.length)];

      //   console.log(history.length, cur, c, val);

      if (val) {
        history.push({ square: cur, attempts: attempts.concat([ val ]) });
        this.setValue(cur.rIndex, cur.cIndex, val);

        count++;
        rollBack = { square: null, attempts: [] };
      } else {
        let diff = [];
        do {
          rollBack = history.pop();
          //    console.log('rollBack', history.length, rollBack);
          this.clearValue(rollBack.square.rIndex, rollBack.square.cIndex);
          count--;

          diff = arrayDiff(this.candidates[rollBack.square.rIndex][rollBack.square.cIndex], rollBack.attempts);
        } while (history.length > 0 && diff.length === 0);
      }
    }
  }
}

/**
 * @param {character[][]} board
 * @return {void} Do not return anything, modify board in-place instead.
 */
var solveSudoku = function(board) {
  const sudoku = new Sudoku(board);
  sudoku.solve();
  board = sudoku.board;
  return sudoku.board;
};

module.exports = { solveSudoku };
