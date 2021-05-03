// interface Laziness {
//   sleep: (time: number) => Laziness
//   sleepFirst: (time: number) => Laziness
//   eat: (food: string) => Laziness
// }

/**
 * @param {string} name
 * @param {(log: string) => void} logFn
 * @returns {Laziness}
 */
function LazyMan(name, logFn) {
  const actions = [[null, `Hi, I'm ${name}.`]];

  const next = () => {
    if (!actions.length) {
      return
    }

    if (actions[0][0] === null) {
      logFn(actions[0][1])
      actions.shift()
      next()
      return
    }
    
    setTimeout(() => {
      actions[0][0] = null
      next()
    }, 1000*actions[0][0])
  }

  setTimeout(next, 0)

  const res = {
    eat: (food) => {
      actions.push([null, `Eat ${food}.`]);

      return res;
    },

    sleep: (time) => {
      const s = time > 1?'s':'';
      actions.push([time, `Wake up after ${time} second${s}.`]);

      return res;
    },

    sleepFirst: (time) => {
      const s = time > 1?'s':'';
      actions.unshift([time, `Wake up after ${time} second${s}.`]);
      return res;
    },
  };

  return res;
}
