module.exports.createRateLimiter = (limit, durationMs) =>
  new RateLimiter(limit, durationMs);

class RateLimiter {
  #logs = [];
  #durationMs = 0;
  #limit = 0;

  constructor(limit, durationMs) {
    this.#limit = limit;
    this.#durationMs = durationMs;
  }

  addRequest() {
    const timeStamp = Date.now();
    this.#logs.push(timeStamp);

    while (
      this.#logs.length > 0 &&
      timeStamp - this.#logs[0] >= this.#durationMs
    ) {
      this.#logs.shift();
    }
  }

  limitExceeded() {
    return this.#logs.length > this.#limit;
  }
}
