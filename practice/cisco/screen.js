/**
 * Given a bowling game outcome, score the game
 * Input: array of 10 items representing the score for each frame.
 * Each frame score is an array consisting of 1-3 items
 */
 const calculateScore = (array) => {
  let c = 0
  let scores = []

  const a = array.flat()
  let j = 0

  for (const frame of array) {
    let score = 0
    for (const n of frame) {
      score += n
      j++
    }

    if (score === 10) {
      c++ // spare

      if (frame.length === 1) {
        c++ //strike
      }
    }

    for (let k = 0; k < c; k++) {
      score += a[j + k] || 0
    }
    c = 0

    scores.push(score)
  }

  let result = 0
  for (const score of scores) {
    result += score
  }
  return result;
}

class TestData {
  constructor(scores, expectedScore) {
    this.scores = scores;
    this.expectedScore = expectedScore;
  }
}

let testCases = [];
const game1 = new TestData(
  [
    [2,3], // 5
    [4,1], // 5
    [5,2], // 7
    [6,2], // 8
    [7,1], // 8
    [7,0], // 7
    [8,1], // 9
    [2,5], // 7
    [1,3], // 4
    [6,2], // 8
  ],
  5 + 5 + 7 + 8 + 8 + 7 + 9 + 7 + 4 + 8
);
testCases.push(game1);

const game2 = new TestData(
  [
    [2, 3], // 5
    [4, 6], // 15 spare
    [5, 2], // 7
    [6, 2], // 8
    [7, 1], // 8
    [7, 0], // 7
    [8, 1], // 9
    [2, 5], // 7
    [1, 3], // 4
    [6, 2], // 8
  ],
  5 + 15 + 7 + 8 + 8 + 7 + 9 + 7 + 4 + 8
);
testCases.push(game2);

const game3 = new TestData(
  [
    [2, 3],  // 5
    [10],    // 17 strike
    [5, 2],  // 7
    [6, 2],  // 8
    [7, 1],  // 8
    [7, 0],  // 7
    [8, 1],  // 9
    [2, 5],  // 7
    [1, 3],  // 4
    [6, 2],  // 8
  ],
  5 + 17 + 7 + 8 + 8 + 7 + 9 + 7 + 4 + 8
);
testCases.push(game3);

const game4 = new TestData(
  [
    [2, 3], // 5
    [4, 6], // 15 spare
    [5, 2], // 7
    [10],   // 18 strike
    [7, 1], // 8
    [7, 3], // 20 spare
    [10],   // 17 strike
    [2, 5], // 7
    [10],   // 18 strike
    [6, 2], // 8
  ],
  5 + 15 + 7 + 18 + 8 + 20 + 17 + 7 + 18 + 8
);
testCases.push(game4);

const game5 = new TestData(
  [
    [2, 3], // 5
    [4, 6], // 20 spare
    [10],   // 18 strike
    [7, 1], // 8
    [10],   // 20 strike
    [7, 3], // 12 spare
    [2, 5], // 7
    [10],   // 26 strike
    [10],   // 18 strike
    [6, 2], // 8
  ],
  5 + 20 + 18 + 8 + 20 + 12 + 7 + 26 + 18 + 8
);
testCases.push(game5);

const game6 = new TestData(
  [
    [2, 3],     // 5
    [4, 6],     // 20
    [10],       // 18
    [7, 1],     // 8
    [10],       // 20
    [7, 3],     // 12
    [2, 5],     // 7
    [10],       // 30 strike
    [10],       // 26 strike
    [10, 6, 4], // 20 spare
  ],
  5 + 20 + 18 + 8 + 20 + 12 + 7 + 30 + 26 + 20
);
testCases.push(game6);

const game7 = new TestData(
  [
    [10],
    [10],
    [10],
    [10],
    [10],
    [10],
    [10],
    [10],
    [10],
    [10, 10, 10]
  ],
  300
);
testCases.push(game7);

///testCases = [game5]

// Print data
testCases.forEach((testCase) => {
  const actual = calculateScore(testCase.scores);
  const expected = testCase.expectedScore;
  const correct = actual == expected;
  const icon = correct ? "\u2705" : "\u274c";

  console.log(icon + "  " + actual + " " + expected);
})
