/**
 * @param {character[][]} board
 * @return {number}
 */
var numRookCaptures = function(board) {
  const ROOK = {}
  let pawn = 'p'

  for (let i = 0; i < board.length; i++) {
    for (let j = 0; j < board[i].length; j++) {
      if (board[i][j] === 'r' || board[i][j] === 'R') {
        ROOK.x = i
        ROOK.y = j

        if (board[i][j] === 'r') {
          pawn = 'P'
        }
      }
    }
  }

  let result = 0

  let x = ROOK.x - 1
  while (x >= 0 && board[x][ROOK.y] === '.') {
    x--
  }
  if (x >= 0 && board[x][ROOK.y] === pawn) {
    result++
  }

  x = ROOK.x + 1
  while (x < 8 && board[x][ROOK.y] === '.') {
    x++
  }
  if (x < 8 && board[x][ROOK.y] === pawn) {
    result++
  }

  let y = ROOK.y - 1
  while (y >= 0 && board[ROOK.x][y] === '.') {
    y--
  }
  if (y >= 0 && board[ROOK.x][y] === pawn) {
    result++
  }

  y = ROOK.y + 1
  while (y < 8 && board[ROOK.x][y] === '.') {
    y++
  }
  if (y < 8 && board[ROOK.x][y] === pawn) {
    result++
  }

  return result
}

const t0 = numRookCaptures([
  [ '.', '.', '.', '.', '.', '.', '.', '.' ],
  [ '.', '.', '.', 'p', '.', '.', '.', '.' ],
  [ 'p', '.', '.', 'R', '.', '.', '.', 'p' ],
  [ '.', '.', '.', '.', '.', '.', '.', '.' ],
  [ '.', '.', '.', '.', '.', '.', '.', '.' ],
  [ '.', '.', '.', 'p', '.', '.', '.', '.' ],
  [ '.', '.', '.', '.', '.', '.', '.', '.' ],
  [ '.', '.', '.', '.', '.', '.', '.', '.' ]
])

console.log(t0)

const t1 = numRookCaptures([
  [ '.', '.', '.', '.', '.', '.', '.', '.' ],
  [ '.', '.', '.', '.', '.', '.', '.', '.' ],
  [ '.', '.', '.', '.', '.', '.', '.', '.' ],
  [ '.', '.', '.', 'R', '.', '.', '.', '.' ],
  [ '.', '.', '.', '.', '.', '.', '.', '.' ],
  [ '.', '.', '.', '.', '.', '.', '.', '.' ],
  [ '.', '.', '.', '.', '.', '.', '.', '.' ],
  [ '.', '.', '.', '.', '.', '.', '.', '.' ]
])

console.log(t1)
