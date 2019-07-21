/**
 * @param {string[]} A
 * @return {number}
 */
var minDeletionSize = function(A) {
  const R = new Set()
  const solve = (arr, firstIndex) => {
    console.log(arr, [ ...R.values() ])

    for (let i = firstIndex; i < arr[0].length; i++) {
      if (R.has(i)) {
        continue
      }

      let ch = ''
      let sorted = true
      let groups = new Map()

      for (let j = 0; j < arr.length; j++) {
        const str = arr[j]
        sorted = str[i] >= ch
        if (str[i] == ch) {
          let set = groups.get(ch) || new Set()
          set.add(j - 1)
          set.add(j)
          groups.set(ch, set)
        }
        if (!sorted) {
          R.add(i)
          break
        }
        ch = str[i]
      }
      if (sorted) {
        for (const equals of groups.values()) {
          solve([ ...equals.values() ].map((index) => arr[index]), i + 1)
        }
        break
      }
    }
  }

  solve(A, 0)
  return R.size
}
/* 
console.log(minDeletionSize([ 'ca', 'bb', 'ac' ]))
console.log(minDeletionSize([ 'xc', 'yb', 'za' ]))
console.log(minDeletionSize([ 'zyx', 'wvu', 'tsr' ]))
console.log(minDeletionSize([ 'xga', 'xfb', 'yfa' ]))

console.log(minDeletionSize([ 'abx', 'agz', 'bgc', 'bfc' ]))

console.log(
  minDeletionSize([
    'bwwdyeyfhc',
    'bchpphbtkh',
    'hmpudwfkpw',
    'lqeoyqkqwe',
    'riobghmpaa',
    'stbheblgao',
    'snlaewujlc',
    'tqlzolljas',
    'twdkexzvfx',
    'wacnnhjdis'
  ])
) */

console.log(
  minDeletionSize([
    'hdbbaomiyk',
    'amcdtrnhjn',
    'fheqnqdkjq',
    'mfeluiclbm',
    'jkexmcstwn',
    'egfmxwjxdj',
    'ayhowbifcx',
    'swhykufgfk',
    'vxhdwxuhwj',
    'johfdcfojv',
    'rnircklfcm',
    'lzkwfqomcz',
    'fvkkhzomgb',
    'aukuoedptv',
    'eimzwmlgxj',
    'ptmnmgppso',
    'oknfgdtweb',
    'mtnukewwir',
    'nlowbhwjdm',
    'tcovbbvuuw',
    'ilqyvtgnfv',
    'nrqgupdyyg',
    'wnrdwmsnzt',
    'rosqrtdeus',
    'bysheeghqg',
    'ciswvgqqlf',
    'uwteztkmqf',
    'tbumqubzdb',
    'dqxbfiwuvm',
    'atxbvdiywo'
  ])
)
