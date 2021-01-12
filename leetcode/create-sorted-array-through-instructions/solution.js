/**
 * @param {number[]} instructions
 * @return {number}
 */
var createSortedArray = function(instructions) {
  const a = []
  let res = 0
  const count = new Map()
  for (const n of instructions) {
    const c = count.get(n) || 0
    if (!a.length || a[a.length-1] <= n) {
      a.push(n)
      count.set(n, c + 1)
      continue
    }
    
    if (a[0] >= n) {
      a.unshift(n)
      count.set(n, c + 1)
      continue
    }
    
    let left = 0
    let l = 0
    let r = a.length
    
    while (l < r) {
      const mid = (l + r) >> 1
      if (a[mid - 1] < n && a[mid] >= n) {
        left = mid
        break
      }
      
      if (a[mid] < n && a[mid+1] >= n) {
        left = mid+1
        break
      }
      
      if (a[mid] >= n) {
        r = mid
        continue
      }
      
      l = mid + 1
    }
    const right = left + c

    res += Math.min(left, a.length - right)
    a.splice(left,0,n)
    count.set(n, c + 1)
  }
  
  return res
};

console.log(createSortedArray([1,2,3,6,5,4]))
console.log(createSortedArray([1,3,3,3,2,4,2,1,2]))
