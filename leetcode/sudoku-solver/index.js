const arrayDiff = function(a, b) {
  return a.filter(function(i) {
    return b.indexOf(i) === -1
  })
}

class Sudoku {
  constructor(board) {
    this.board = board
    this.indices = board.map((line, i) => i)
    this.digits = this.indices.map((i) => (i + 1).toString())
  }

  getColumn(index) {
    return this.board.map((line) => line[index])
  }

  getColumnCandidates(index) {
    return arrayDiff(this.digits, this.getColumn(index))
  }

  generateCandidates() {
    const columnCandidates = this.indices.map((i) => this.getColumnCandidates(i))

    return this.board.map((row, i) => row.map((cell, j) => arrayDiff(columnCandidates[j], row)))
  }

  get candidates() {
    this._candidates = this._candidates || this.generateCandidates()
    return this._candidates
  }

  resetCaches() {
    this._candidates = null
  }

  isSquareEmpty(rIndex, cIndex) {
    return this.board[rIndex][cIndex] === '.'
  }

  setValue(rIndex, cIndex, val) {
    this.board[rIndex][cIndex] = val
    this.indices.map((i) => {
      this._candidates[i][cIndex] = this._candidates[i][cIndex].filter((v) => val !== v)
      this._candidates[rIndex][i] = this._candidates[rIndex][i].filter((v) => val !== v)
    })
  }

  clearValue(rIndex, cIndex) {
    const val = this.board[rIndex][cIndex]
    this.board[rIndex][cIndex] = '.'
    this.indices.map((i) => {
      this._candidates[i][cIndex].push(val)
      this._candidates[rIndex][i].push(val)
    })
  }

  solve() {
    const emptySquares = this.indices.reduce((acc, cur, i) => {
      this.indices.map((j) => {
        if (this.isSquareEmpty(i, j)) {
          acc.push({
            rIndex: i,
            cIndex: j,
            priority: 1
            // tried: [],
            // candidates: this.candidates[i][j],
            // weight: this.candidates[i][j].length
          })
        }
      })

      return acc
    }, [])

    const getSmallest = () => {
      return emptySquares.reduce(
        (lowest, next, index) =>
          next.priority <= lowest.priority &&
          this.candidates[next.rIndex][next.cIndex].length > 0 &&
          this.candidates[next.rIndex][next.cIndex].length <
            /*        */ this.candidates[lowest.rIndex][lowest.cIndex].length
            ? next
            : lowest,
        emptySquares[0]
      )
    }
    const stack = []
    let count = 0
    while (emptySquares.length > count) {
      const cur = getSmallest()
      console.log(stack.length, cur)

      const c = this.candidates[cur.rIndex][cur.cIndex]
      const val = c[0] //c[Math.floor(Math.random() * c.length)]
      if (val) {
        stack.push(cur)
        cur.priority++
        this.setValue(cur.rIndex, cur.cIndex, val)

        count++
      } else {
        let rollBack
        do {
          rollBack = stack.pop()
          console.log('rollBack', rollBack)
          this.clearValue(rollBack.rIndex, rollBack.cIndex)
          count--
        } while (this.candidates[rollBack.rIndex][rollBack.cIndex].length < 3)
      }
    }
  }
}

/**
 * @param {character[][]} board
 * @return {void} Do not return anything, modify board in-place instead.
 */
var solveSudoku = function(board) {
  const sudoku = new Sudoku(board)
  sudoku.solve()
  return sudoku.board
}

module.exports = { solveSudoku }
