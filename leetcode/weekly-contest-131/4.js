/**
 * @param {number[][]} clips
 * @param {number} T
 * @return {number}
 */
var videoStitching = function(clips, T) {
  clips = clips.sort((a, b) => {
    if (a[0] === b[0]) {
      return b[1] - a[1]
    }

    return a[0] - b[0]
  })

  //console.log(clips)

  /*  if (clips[0][0] > 0 || clips[clips.length - 1][1] < T) {
    return -1
  } */

  let [ r, i, t ] = [ 0, 0, 0 ]

  while (i < clips.length && t < T) {
    let m = 0
    while (i < clips.length && clips[i][0] <= t) {
      m = Math.max(m, clips[i][1])
      i++
    }
    if (m == 0) {
      break
    }

    r++
    t = m
  }

  if (t < T) {
    r = -1
  }

  return r
}

// prettier-ignore
console.log(videoStitching([[0,1],[6,8],[0,2],[5,6],[0,4],[0,3],[6,7],[1,3],[4,7],[1,4],[2,5],[2,6],[3,4],[4,5],[5,7],[6,9]],
  9))

/* 
console.log(videoStitching([ [ 0, 2 ], [ 0, 3 ], [ 3, 8 ] ], 5))

console.log(videoStitching([ [ 0, 2 ], [ 4, 6 ], [ 8, 10 ], [ 1, 9 ], [ 1, 5 ], [ 5, 9 ] ], 10))
console.log(videoStitching([ [ 0, 1 ], [ 1, 2 ] ], 5))
console.log(
  videoStitching(
    [
      [ 0, 1 ],
      [ 6, 8 ],
      [ 0, 2 ],
      [ 5, 6 ],
      [ 0, 4 ],
      [ 0, 3 ],
      [ 6, 7 ],
      [ 1, 3 ],
      [ 4, 7 ],
      [ 1, 4 ],
      [ 2, 5 ],
      [ 2, 6 ],
      [ 3, 4 ],
      [ 4, 5 ],
      [ 5, 7 ],
      [ 6, 9 ]
    ],
    9
  )
)
 */
