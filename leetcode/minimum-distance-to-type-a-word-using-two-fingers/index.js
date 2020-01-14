/**
 * @param {string} word
 * @return {number}
 */
const minimumDistance = (word) => {
  const cMemo = {};
  const getCoords = (ch) => {
    if (ch in cMemo) {
      return cMemo[ch];
    }
    const ind = ch.length === 1 ? ch.charCodeAt(0) - 'A'.charCodeAt(0) : 0;
    cMemo[ch] = [ Math.floor(ind / 6), ind % 6 ];

    return cMemo[ch];
  };

  const memo = word.split('').map(() => ({}));

  const getDist = (p1, p2) => {
    return Math.abs(p1[0] - p2[0]) + Math.abs(p1[1] - p2[1]);
  };

  const getKey = (state) => {
    return `${state.l}-${state.r}`;
  };

  const states = [ { pos: 0, l: '', r: '', cost: 0 } ];

  let r = states.length ? Number.POSITIVE_INFINITY : 0;

  while (states.length) {
    const state = states.pop();
    const key = getKey(state);

    if (key in memo[state.pos] && state.cost > memo[state.pos][key]) {
      continue;
    }

    const ch = word[state.pos];
    const coord = getCoords(ch);

    let dl = 0;
    if (state.l !== '') {
      dl = getDist(getCoords(state.l), coord);
    }

    let dr = 0;
    if (state.r !== '') {
      dr = getDist(getCoords(state.r), coord);
    }

    if (state.pos < word.length - 1) {
      const left = { r: state.r, pos: state.pos + 1, l: ch, cost: state.cost + dl };
      const right = state;
      right.pos++;
      right.r = ch;
      right.cost += dr;

      const [ lKey, rKey ] = [ getKey(left), getKey(right) ];

      if (!(lKey in memo[left.pos]) || left.cost < memo[left.pos][lKey]) {
        memo[left.pos][lKey] = left.cost;
        states.push(left);
      }

      if (!(rKey in memo[right.pos]) || right.cost < memo[right.pos][rKey]) {
        memo[right.pos][rKey] = right.cost;
        states.push(right);
      }
    } else {
      r = Math.min(r, state.cost + dl, state.cost + dr);
    }
  }

  return r;
};
/* 
console.log(minimumDistance('CAKE'));
console.log(minimumDistance('HAPPY'));
console.log(minimumDistance('NEW'));
console.log(minimumDistance('YEAR'));
console.log(
  minimumDistance(
    'OPVUWZLCKTDPSUKGHAXIDWHLZFKNBDZEWHBSURTVCADUGTSDMCLDBTAGFWDPGXZBVARNTDICHCUJLNFBQOBTDWMGILXPSFWVGYBZVFFKQIDTOVFAPVNSQJULMVIERWAOXCKXBRI'
  )
);
 */
/* console.log(
  minimumDistance(
    'EKUSSPCQPNKMXADNTVEYEADWNHBUIIMDLJLTZUGXRRBKQRAFPAUIKELDTRNZYCISDQCSEPBQZHLDRWGFWVHMGTWXOLFIXLRGGOKEIRVYY'
  )
); */

console.log(
  minimumDistance(
    'YMFEVNMOXXUXGICWFSAJXFDEIFMIQUNKRXNPODITKIIXJMTQAPUQQGNIXUYNXBANRCZIWZAAPVZMRLRPQORXXWPSFZSYATLGJNCHDFZTOUEPKOWEPZIPKWQXZHYBLQIMRUTDJEQEECKGBAPPCSQBEUOSESICJLEOLBKDDTVDJVJSLFYAVYPMMUGCNIPNVPMENZGYNUYGICZEOKLJZZOKHFOEWKZHPEDRSTOXKFCVKFILVWDDMKYKJKYAGVUJUFFRFPHCWQIWDBMXCMOPVEWE'
  )
);
