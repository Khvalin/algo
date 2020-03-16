
function program(robot) {
   const scanner = robot.getScanner();
   
  const getDirs = (p1, p2) => {
    res = []
    if (p1.getX() < p2.getX()) {
      res.push('RIGHT')
    }
    if (p1.getX() > p2.getX()) {
      res.push('LEFT')
    }
    
    if (p1.getY() < p2.getY()) {
      res.push('UP')
    }
    if (p1.getY() > p2.getY()) {
      res.push('DOWN')
    }

    return res
  }
  
  const to_key = (p) => {
    return p.getX() + ' ' + p.getY()
  }
  
  const dfs = (start, end) => {
    const pathLen = {}
    
    const stack = [{point: start, path:[] }];
    pathLen[to_key(start)] = 0

    while (stack.length > 0 ) {
      const { point, path } = stack.shift()
      
      const len = pathLen[to_key(point)]
      
      if (to_key(point) === to_key(end)) {
        return path
      }
      
      for (let i = -1; i <= 1; i++) {
        for (let j = -1; j <= 1; j++) {
          if (i * j != 0 || i === j) {
            continue
          }
          
          let [x, y] = [point.getX() + i, point.getY() + j]
          if ( scanner.isAnyOfAt(x, y, ["HOLE", "WALL"]) ) {
            continue
          }
           
          
          const p = new Point(x,y)
            
          pathLen[to_key(p)] = pathLen[to_key(p)] || Number.POSITIVE_INFINITY
            
          if (pathLen[to_key(p)] <= len + 1){
            continue
          }
          
          robot.log(x + ' ' + y + ' ' + scanner.getAt(x,y))
            
          pathLen[to_key(p)] = len + 1
          stack.push({point: p, path: [...path, p]})
        }
      }
    }
    
    return []
  }
  
  const me = scanner.getMe()
  const memory = robot.getMemory();
  
  const createPath = () => {
    let golds = scanner.getGold();
    if (golds.length > 0) {
      return dfs(me, golds[0])
    }
  }
  
  let path = memory.load("PATH") || []
  
  
  if (!path.length) {
    path = createPath()
  }
   robot.log(me)
   robot.log(path)
  
  const next = path.shift()
 
  
  robot.go(getDirs(me, next)[0])

   
  memory.save("PATH", path)
}