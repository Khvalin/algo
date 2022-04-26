/**
 * @param {number[][]} points
 * @return {number}
 */
var minCostConnectPoints = function (points) {
    const parent = points.map((_, i) => i)

    const find = (i) => {
        while (i != parent[i]) {
            i = parent[i];
        }
        return i;
    }

    const uni = (i, j) => {
        if (i != j) {
            parent[j] = i;
            return 1;
        }
        return 0;
    }

    const edges = []

    for (let i = 0; i < points.length; i++) {
        for (let j = i + 1; j < points.length; j++) {
            const d = Math.abs(points[i][0] - points[j][0]) + Math.abs(points[i][1] - points[j][1]);
            edges.push([d, i, j])
        }
    }

    edges.sort((a, b) => a[0] - b[0])
    console.log(edges)

    let count = 0
    let i = 0

    let res = 0

    while (count < points.length - 1 && i < edges.length) {
        const [d, u, v] = edges[i]
        i++

        u = find(u);
        v = find(v);
        if (v === u) {
            continue
        }

        uni(v, u)

        res += d
        count++
    }

    return res
};