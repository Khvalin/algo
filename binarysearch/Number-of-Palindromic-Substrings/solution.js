class Solution {
  solve(s) {
    let res = 0;

    const lookUp = (index) => {
      for (let j = 0; j < 2; j++) {
        let i = 0;
        while (index - i >= 0 && index + i + j < s.length && s[index - i] === s[index + i + j]) {
          res++;
          i++;
        }
      }
    };

    for (var i = 0; i < s.length; i++) {
      lookUp(i);
    }

    return res;
  }
}
