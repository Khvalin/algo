/**
 * @param {number[]} commands
 * @param {number[][]} obstacles
 * @return {number}
 */
var robotSim = function(commands, obstacles) {
  const UP = { x: 0, y: 1 }
  const DOWN = { x: 0, y: -1 }
  const LEFT = { x: -1, y: 0 }
  const RIGHT = { x: 1, y: 0 }

  const DIRECTIONS = [ UP, RIGHT, DOWN, LEFT ]

  const obstaclesMap = {}

  for (obs of obstacles) {
    obstaclesMap[obs.join(' ')] = true
  }

  //console.log(obstaclesMap)

  let curPos = { x: 0, y: 0 }
  let dirIndex = 0
  let result = 0

  for (const command of commands) {
    if (command === -2) {
      dirIndex -= 1
      if (!DIRECTIONS[dirIndex]) {
        dirIndex = DIRECTIONS.length - 1
      }
    }
    if (command === -1) {
      dirIndex += 1
      if (!DIRECTIONS[dirIndex]) {
        dirIndex = 0
      }
    }

    if (command > 0) {
      const dir = DIRECTIONS[dirIndex]
      for (let i = 0; i < command; i++) {
        //  console.log(dir)

        const nextPos = { x: curPos.x + dir.x, y: curPos.y + dir.y }

        if (!obstaclesMap[`${nextPos.x} ${nextPos.y}`]) {
          curPos.x = nextPos.x
          curPos.y = nextPos.y
          result = Math.max(result, curPos.x * curPos.x + curPos.y * curPos.y)
        } else {
          break
        }

        // console.log(curPos)
      }
    }
  }

  // console.log(curPos)
  return result
}

console.log(robotSim([ 4, -1, 3 ], []))
console.log(robotSim([ 4, -1, 4, -2, 4 ], [ [ 2, 4 ] ]))
console.log(
  robotSim(
    [ -2, -1, 8, 9, 6 ],
    [ [ -1, 3 ], [ 0, 1 ], [ -1, 5 ], [ -2, -4 ], [ 5, 4 ], [ -2, -3 ], [ 5, -1 ], [ 1, -1 ], [ 5, 5 ], [ 5, 2 ] ]
  )
)

console.log(
  robotSim(
    [
      2,
      5,
      2,
      5,
      -1,
      3,
      4,
      2,
      4,
      -2,
      5,
      8,
      -1,
      -2,
      -2,
      -1,
      8,
      -1,
      -2,
      -2,
      -1,
      -1,
      5,
      -1,
      -1,
      1,
      5,
      9,
      1,
      -1,
      -2,
      -1,
      -2,
      3,
      7,
      2,
      3,
      5,
      9,
      -2,
      5,
      2,
      1,
      4,
      6,
      5,
      9,
      1,
      9,
      6,
      3,
      -1,
      -1,
      9,
      9,
      -1,
      1,
      6,
      3,
      -2,
      2,
      2,
      5,
      -1,
      -1,
      -2,
      9,
      6,
      5,
      5,
      9,
      2,
      5,
      -2,
      -2,
      5,
      7,
      -1,
      -2,
      -1,
      2,
      -1,
      1,
      -1,
      -1,
      2,
      8,
      -1,
      8,
      3,
      -2,
      3,
      -1,
      -2,
      4,
      7,
      3,
      8,
      -1,
      2
    ],
    [
      [ -5, -77 ],
      [ 35, 40 ],
      [ 58, -30 ],
      [ 31, -96 ],
      [ 40, 14 ],
      [ -25, 50 ],
      [ 37, -38 ],
      [ -54, -31 ],
      [ 64, -41 ],
      [ 72, 53 ],
      [ 83, -95 ],
      [ -31, -61 ],
      [ 68, 32 ],
      [ -56, 16 ],
      [ 88, -81 ],
      [ -48, -31 ],
      [ 56, -57 ],
      [ -74, 24 ],
      [ -42, -59 ],
      [ 72, -86 ],
      [ 40, 34 ],
      [ -85, -45 ],
      [ 22, -35 ],
      [ -95, 56 ],
      [ -66, 42 ],
      [ -67, 94 ],
      [ 46, 10 ],
      [ 35, 27 ],
      [ -9, -6 ],
      [ -84, -97 ],
      [ 38, -22 ],
      [ 3, -39 ],
      [ 71, -97 ],
      [ -40, -86 ],
      [ -45, 56 ],
      [ -91, 59 ],
      [ 24, -11 ],
      [ 91, 100 ],
      [ -68, 43 ],
      [ 34, 27 ],
      [ -60, 32 ],
      [ -20, 34 ],
      [ -34, 34 ],
      [ -31, -53 ],
      [ 52, -98 ],
      [ -70, -15 ],
      [ 73, -41 ],
      [ -94, 95 ],
      [ -78, -42 ],
      [ -7, -11 ],
      [ -37, -94 ],
      [ 26, -74 ],
      [ -53, 68 ],
      [ 72, 2 ],
      [ -80, -58 ],
      [ -94, 48 ],
      [ -80, -57 ],
      [ -35, 69 ],
      [ 17, -45 ],
      [ -72, -76 ],
      [ 21, 99 ],
      [ -25, -19 ],
      [ -48, 86 ],
      [ 86, -47 ],
      [ 59, -66 ],
      [ -38, -16 ],
      [ 47, -44 ],
      [ 28, 96 ],
      [ 92, -64 ],
      [ -62, -35 ],
      [ -63, -87 ],
      [ -83, 95 ],
      [ 25, -89 ],
      [ 30, -40 ],
      [ -75, 93 ],
      [ 47, -21 ],
      [ 12, -93 ],
      [ 70, -22 ],
      [ -64, -43 ],
      [ -17, -86 ],
      [ -33, 93 ],
      [ -74, -7 ],
      [ -78, 5 ],
      [ -37, -11 ],
      [ -84, -29 ],
      [ -29, -11 ],
      [ 17, 9 ],
      [ -64, 10 ],
      [ 25, 29 ],
      [ 14, 25 ],
      [ 19, 42 ],
      [ 71, 52 ],
      [ 30, -76 ],
      [ 19, 66 ],
      [ 40, 99 ],
      [ -61, -95 ],
      [ -17, 40 ],
      [ 6, -21 ],
      [ 7, 28 ],
      [ -4, 85 ],
      [ 71, 99 ],
      [ 50, 99 ],
      [ -53, -95 ],
      [ -8, 8 ],
      [ 63, -79 ],
      [ 88, -35 ],
      [ 50, -38 ],
      [ -60, -31 ],
      [ -2, -8 ],
      [ -8, 91 ],
      [ -14, 50 ],
      [ -25, -26 ],
      [ 1, 71 ],
      [ -91, -64 ],
      [ -40, 46 ],
      [ 30, -97 ],
      [ 9, -55 ],
      [ -36, -75 ],
      [ 71, 99 ],
      [ 90, -53 ],
      [ -68, -20 ],
      [ -73, 89 ],
      [ -14, 74 ],
      [ -8, 72 ],
      [ 82, 48 ],
      [ 45, 2 ],
      [ -42, 12 ],
      [ 77, 22 ],
      [ 87, 56 ],
      [ 73, -21 ],
      [ 78, 15 ],
      [ -6, -75 ],
      [ 24, 46 ],
      [ -11, -70 ],
      [ -90, -77 ],
      [ 57, 43 ],
      [ -66, 10 ],
      [ -30, -47 ],
      [ 91, -37 ],
      [ -4, -67 ],
      [ -88, 19 ],
      [ 66, 29 ],
      [ 86, 97 ],
      [ -4, 67 ],
      [ 54, -92 ],
      [ -54, 71 ],
      [ -42, -17 ],
      [ 57, -91 ],
      [ -9, -15 ],
      [ -26, 56 ],
      [ -57, -58 ],
      [ -72, 91 ],
      [ -55, 35 ],
      [ -20, 29 ],
      [ 51, 70 ],
      [ -61, 88 ],
      [ -62, 99 ],
      [ 52, 17 ],
      [ -75, -32 ],
      [ 91, -22 ],
      [ 54, 33 ],
      [ -45, -59 ],
      [ 47, -48 ],
      [ 53, -98 ],
      [ -91, 83 ],
      [ 81, 12 ],
      [ -34, -90 ],
      [ -79, -82 ],
      [ -15, -86 ],
      [ -24, 66 ],
      [ -35, 35 ],
      [ 3, 31 ],
      [ 87, 93 ],
      [ 2, -19 ],
      [ 87, -93 ],
      [ 24, -10 ],
      [ 84, -53 ],
      [ 86, 87 ],
      [ -88, -18 ],
      [ -51, 89 ],
      [ 96, 66 ],
      [ -77, -94 ],
      [ -39, -1 ],
      [ 89, 51 ],
      [ -23, -72 ],
      [ 27, 24 ],
      [ 53, -80 ],
      [ 52, -33 ],
      [ 32, 4 ],
      [ 78, -55 ],
      [ -25, 18 ],
      [ -23, 47 ],
      [ 79, -5 ],
      [ -23, -22 ],
      [ 14, -25 ],
      [ -11, 69 ],
      [ 63, 36 ],
      [ 35, -99 ],
      [ -24, 82 ],
      [ -29, -98 ]
    ]
  )
)
