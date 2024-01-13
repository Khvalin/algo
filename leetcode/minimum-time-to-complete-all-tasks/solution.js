/**
 * @param {number[][]} tasks
 * @return {number}
 */
var findMinimumTime = function(tasks) {   
    const starts = new Map()
    const ends = new Map()
    
    const durations = []
    
    let min = tasks[0][0]
    let max = tasks[0][1]
    for (let i = 0; i < tasks.length; i++) {
        const [s, e, d] = tasks[i]
        durations.push(d)
        
        let a = starts.get(s) || []
        a.push(i)
        starts.set(s, [a])
        
        a = ends.get(e) || []
        a.push(i)
        ends.set(e, [a])
        
        min = Math.min(s, min)
        max = Math.max(e, max)
    }
    
    let res = 0
    
    const scheduled = new Set()
    for (let t = min; t <= max; t++) {
        // console.log(durations)
        let a = ends.get(t) || []
        let maxDuration = 0
        for (const ind of a) {
            maxDuration = Math.max(maxDuration, durations[ind])
            scheduled.delete(ind)
        }


        if (maxDuration > 0) {
            console.log(t, maxDuration)
            res -= maxDuration
            for (const ind of scheduled.values()) {
               durations[ind] -= Math.min(maxDuration, t - tasks[ind] + 1)
            }
        }     
        
        a = starts.get(t) || []
        for (const ind of a) {
            scheduled.add(ind)
        }
    }
    
    return max - min + res +1
};


let res
res = findMinimumTime([[2,3,1],[4,5,1],[1,5,2]])
console.log(res)

res = findMinimumTime([[1,3,2],[2,5,3],[5,6,2]])
console.log(res)

