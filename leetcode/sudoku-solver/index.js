class Sudoku {
  constructor(board) {
    this.board = board.map((line) => line.map((cell) => ('.' === cell ? 0 : 1 * cell)))
  }
}

/**
 * @param {character[][]} board
 * @return {void} Do not return anything, modify board in-place instead.
 */
var solveSudoku = function(board) {
  const sudoku = new Sudoku(board)
  return sudoku.board
}

module.exports = { solveSudoku }
