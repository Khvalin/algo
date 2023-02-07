module.exports = {
  //param A : array of integers
  //return a array of array of integers
  permute: function (A) {
    const res = [];
    const q = [[[], A]];

    while (q.length) {
      const [p, a] = q.pop();
      if (!a.length) {
        res.push(p);
        continue;
      }

      for (let i = 0; i < a.length; i++) {
        const pCopy = [...p, a[i]];
        const aCopy = [...a];
        aCopy.splice(i, 1);
        q.push([pCopy, aCopy]);
      }
    }

    return res;
  },
};
