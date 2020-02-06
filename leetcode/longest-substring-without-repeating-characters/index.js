var lengthOfLongestSubstring = function(s) {
  const m = new Map()
  let [r, p] = [0, -1]

  for (let i = 0; i < s.length; i++){
    const ch = s[i]

    if (m.has(ch)) {
      p = Math.max(p, m.get(ch))
    }
    r = Math.max(r, i - p)

    m.set(ch, i)
  }
  
  return r
};

console.log(lengthOfLongestSubstring("aazzabcccb"))
