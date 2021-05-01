/**
 * @param {string[]} deadends
 * @param {string} target
 * @return {number}
 */
var openLock = function(des, target) {
  const toArray = (n) => {
    let r = []
    for (let i = 0; i < 4; i++) {
      r.unshift(n%10)
      n = Math.floor(n / 10)
    }
    
    return r
  }
  
  const toNum = (arr) => {
    let r = 0
    for (let i = 0; i< arr.length; i++) {
      r *= 10
      r += arr[i]
    }

    return r
  }

  const deadends = new Set()  
  for (const d of des) {
    let s = 0
    for (let i = 0; i < d.length; i++) {
      s *= 10
      s += d.charCodeAt(i) - '0'.charCodeAt(0)
    }

    deadends.add(s)
  }
  
  let t = 0
  for (let i = 0; i < target.length; i++) {
    t *= 10
    t += target.charCodeAt(i) - '0'.charCodeAt(0)
  }
  
  const dist = [0]
  for (let i = 0; i < 10010; i++) {
    dist.push(Number.POSITIVE_INFINITY)
  }

  let res = Number.POSITIVE_INFINITY
  
  const q = []
  if (!deadends.has(0)) {
    q.push(0)
  }
  while (q.length) {
    let c = q.shift()
    const d = dist[c]
    let cur = toArray(c)
    
    for (let k = -1; k < 2; k += 2) {
      for (let i = 0; i < 4; i++) {
        const next = [...cur]
        next[i] = (next[i] + 10 + k) % 10
        const n = toNum(next)
        if (!deadends.has(n) && dist[n] > d + 1) {
          dist[n] = d + 1
          q.push(n)
        }
      }
    }
  }
  
  return dist[t] < Number.POSITIVE_INFINITY? dist[t]:-1
};

let deadends = ["0201","0101","0102","1212","2002"], target = "0202"
deadends = ["1002","1220","0122","0112","0121"]
target = "1200"

console.log(openLock(deadends, target))
