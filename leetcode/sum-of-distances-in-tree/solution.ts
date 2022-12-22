function sumOfDistancesInTree(n: number, edges: number[][]): number[] {
    const res: number[] = new Array(n);
    const adj = new Array(n+1);
    for (let i = 0; i < n; i++) {
        adj[i] = []
        res[i] = 0
    }

    for (const [u, v] of edges) {
        adj[u].push(v)
        adj[v].push(u)
    }

    const count = (ind) => {
        const seen = new Array(n+1);
        let c = 0
        const q:number[][] = [[ind, 1]]
        seen[ind] = true
        while (q.length) {
            const [i, d] = q.shift()
            for (let j of adj[i]) {
                if (seen[j]) {
                    continue
                }
                seen[j] = true
                q.push([j, d + 1])
                c += d
            }
        }

        return c
    }

    for (let i = 0; i < n; i++) {
        res[i] = count(i)
    }

    return res;
};

let res = sumOfDistancesInTree(6, [[0,1],[0,2],[2,3],[2,4],[2,5]])
console.log(res)