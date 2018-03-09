/**
 * @param {string[]} board
 * @return {boolean}
 */

var validTicTacToe = function(board) {
  const X = 'X';
  const O = 'O';

  const getXOCount = (board) => {
    const s = board.join('');

    return {
      x: (s.match(/X/g) || []).length,
      o: (s.match(/O/g) || []).length
    };
  };

  const getWinState = (board, side) => {
    const triple = side.repeat(3);

    let result = false;
    for (let i = 0; i < board.length; i++) {
      const vertical = board[i][0] === side && board[i][1] === side && board[i][2] === side;
      result = result || board[i] === triple || vertical;
    }

    return (
      result ||
      (board[1][1] === side && board[0][0] === side && board[2][2] === side) ||
      (board[1][1] === side && board[0][2] === side && board[2][0] === side)
    );
  };

  const counts = getXOCount(board);
  const xWin = getWinState(board, X);
  const oWin = getWinState(board, O);

  if (
    (counts.x === 0 && counts.o > 0) ||
    counts.o > counts.x ||
    counts.x > counts.o + 1 ||
    (xWin && counts.x !== counts.o + 1) ||
    (oWin && counts.x !== counts.o)
  ) {
    return false;
  }

  return true;
};

console.log(validTicTacToe([ 'O  ', '   ', '   ' ]));
console.log(validTicTacToe([ 'XOX', ' X ', '   ' ]));
console.log(validTicTacToe([ 'XXX', '   ', 'OOO' ]));
console.log(validTicTacToe([ 'XOX', 'O O', 'XOX' ]));
console.log(validTicTacToe([ 'XOX', 'OOO', ' XX' ]));
