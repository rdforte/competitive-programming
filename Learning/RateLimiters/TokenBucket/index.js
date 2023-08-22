module.exports.createRateLimiter = (capacity, fillPerSecond) =>
  new TokenBucket(capacity, fillPerSecond);

class TokenBucket {
  #capacity;
  #fillPerSecond;
  #lastFilled;
  #tokens;

  constructor(capacity, fillPerSecond) {
    this.#capacity = capacity;
    this.#fillPerSecond = fillPerSecond;

    this.#lastFilled = Date.now();
    this.#tokens = capacity;
  }

  take() {
    this.#refill();
    this.#tokens = Math.max(0, this.#tokens - 1);
  }

  exceededLimit() {
    this.#refill();
    return this.#tokens === 0;
  }

  #refill() {
    const now = Date.now();
    const elapsedTimeMs = now - this.#lastFilled;
    const elapsedTimeSec = elapsedTimeMs / 1000;
    const rate = elapsedTimeSec / this.#fillPerSecond;

    this.#tokens = Math.min(
      this.#capacity,
      this.#tokens + Math.floor(rate * this.#capacity)
    );
    this.#lastFilled = now;
  }
}
