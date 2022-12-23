let startTime = null;

let pushIntervalId = null;

const progressLabel = document.querySelector(".progress");
const totalTimeLabel = document.querySelector(".total-time");

const setProgress = (p) => {
  progressLabel.innerText = p;
};

const setTotalTime = (t) => {
  totalTimeLabel.innerText = t;
};

const messageHandler = (eventData) => {
  const { complete, total } = eventData;
  window.requestAnimationFrame(() => {
    const p = `${(100 * complete / total).toFixed(2)}%`;
    setProgress(p);
  });
};

const doneHandler = (eventData) => {
  const endTime = new Date().getTime();
  setTotalTime(`${(endTime - startTime) / 1000} s`);

  setProgress("done");
  console.log(eventData.array.splice(0, 1500));
  window.clearInterval(pushIntervalId);
};

const handlers = {
  onprogress: messageHandler,
  done: doneHandler,
  NOOP: () => {},
};

const init = () => {
  const sorterWorker = new Worker("js/worker.js");

  sorterWorker.postMessage({ "action": "sort" });

  startTime = new Date().getTime();

  pushIntervalId = window.setInterval(() => {
    const num = Math.floor(Math.random() * 1000);
    sorterWorker.postMessage({ "action": "push", "num": num });
  }, 200);

  sorterWorker.onmessage = (e) => {
    const action = e.data.action || "NOOP";

    handlers[action]?.call(this, e.data);
  };
};

init();
