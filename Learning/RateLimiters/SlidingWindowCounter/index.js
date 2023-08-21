module.exports.createRateLimiter = (maxRequests, interval) =>
  new RateLimiter(maxRequests, interval);

class RateLimiter {
  #maxRequests;
  #interval;
  #windows = [];

  constructor(maxRequests, interval) {
    this.#maxRequests = maxRequests;
    this.#interval = interval;
    const previousInterval = Date.now();
    const currentInterval = previousInterval + interval;

    this.#windows.push({
      count: 0,
      upperTimeRange: previousInterval,
    });

    this.#windows.push({
      count: 0,
      upperTimeRange: currentInterval,
    });
  }

  addRequest() {
    this.#updateWindows();
    const curWindow = this.#windows[1];
    curWindow.count++;
  }

  exceededLimit() {
    this.#updateWindows();
    const curWindow = this.#windows[1];
    const prevWindow = this.#windows[0];
    const totalRequests = curWindow.count + prevWindow.count;
    const overlapPercentage = prevWindow.count / totalRequests;
    const estimatedRequests =
      curWindow.count + prevWindow.count * overlapPercentage;

    return Math.round(estimatedRequests) >= this.#maxRequests;
  }

  #updateWindows() {
    const time = Date.now();
    const prevWindow = this.#windows[0];
    const curWindow = this.#windows[1];

    const timeMoveToNextWindow =
      time >= curWindow.upperTimeRange &&
      time < curWindow.upperTimeRange + this.#interval;

    if (timeMoveToNextWindow) {
      const newCurrentTime = curWindow.upperTimeRange + this.#interval;
      prevWindow.count = curWindow.count;
      prevWindow.upperTimeRange = curWindow.upperTimeRange;

      curWindow.upperTimeRange = newCurrentTime;
      curWindow.count = 0;
    }

    const timeExceededBothWindows =
      time >= curWindow.upperTimeRange + this.#interval * 2;

    if (timeExceededBothWindows) {
      prevWindow.upperTimeRange = time;
      prevWindow.count = 0;

      curWindow.upperTimeRange = time + this.#interval;
      curWindow.count = 0;
    }
  }
}
