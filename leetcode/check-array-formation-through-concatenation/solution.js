/**
 * @param {number[]} arr
 * @param {number[][]} pieces
 * @return {boolean}
 */
var canFormArray = function(arr, pieces) {
    const m = new Map()
    for (let i = 0; i < pieces.length; i++) {
        const a = pieces[i]
        for (const n of a){
            m.set(n, i)
        }
    }

    for (let i = 0; i < arr.length; i++) {
      if (!m.has(arr[i])) {
        return false
      }

      const ind = m.get(arr[i])
      const a = pieces[ind]
      for (let j = 0; j < a.length; j++) {
        if (a[j] === arr[i]) {
          i++
          continue
        }
        
        return false
      }
    }
    
    return true
};
