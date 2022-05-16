/**
 * @param {number[][]} grid
 * @return {number}
 */
var shortestPathBinaryMatrix = function(grid) {
    const N  = grid.length
    const inf = N * N + 10
    for (const row of grid) {
        for (let i = 0 ;i < row.length; i++) {
            row[i] *= -1
            if (!row[i]) {
                row[i] = inf
            }
        }
    }
    
    
    const q = []
    if (grid[0][0] > 0) {
        grid[0][0] = 1
        q.push([0,0, 1])
    }
    
    while (q.length) {
        const c = q.shift()
        
        const [x,y] = c
        
        d = grid[x][y]

        for (let dx = 1; dx >= -1; dx--) { // A* type of an approach
            for (let dy = 1; dy >= -1; dy--) {
                const [x1,y1] = [x+dx, y+dy]
                
                if (x1 < 0 || x1 >= N || y1 < 0 || y1 >= N  || grid[x1][y1] <= d + 1) {
                    continue
                }
                
                q.push([x1,y1])
                grid[x1][y1] = d + 1
                
                if (x1 === N - 1 && y1 === N - 1) {
                    return d + 1
                }
            }        
        }
    }
    
   
    
    return grid[N-1][N-1] === inf? -1:grid[N-1][N-1]
};
