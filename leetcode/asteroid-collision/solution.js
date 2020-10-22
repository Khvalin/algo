/**
 * @param {number[]} asteroids
 * @return {number[]}
 */
var asteroidCollision = function(asteroids) {
    const a = asteroids
    for (let i = 0; i < a.length - 1; ) {
      if (a[i] > 0 && a[i + 1] < 0) {
        const d = Math.abs(a[i] + a[i + 1]) * Math.sign(a[i] + a[i + 1])
        const next = []
        if (d !== 0) {
          next.push( Math.sign(d) * Math.max(Math.abs(a[i]), Math.abs(a[i + 1])) )
        }
        
        a.splice(i, 2, ...next)

        
        i = Math.max(i - 1, 0)
        continue
      }
        
      i++
    }
    
    return a
};

console.log(asteroidCollision([5,10,-5]))
console.log(asteroidCollision([5,-5]))
console.log(asteroidCollision([10,2,-5]))
console.log(asteroidCollision([-2,-1,1,2]))
