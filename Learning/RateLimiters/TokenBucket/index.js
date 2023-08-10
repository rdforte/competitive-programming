module.exports.createRateLimiter = (bucketSize) => new RateLimiter(bucketSize);

class RateLimiter {
  #bucket;
  #timer;
  #maxBucketSize;

  constructor(bucketSize) {
    this.#bucket = bucketSize;
    this.#maxBucketSize = bucketSize;
  }

  addRequest() {
    this.#bucket = Math.max(0, this.#bucket - 1);
  }

  exceededLimit() {
    return this.#bucket === 0;
  }

  startRunRefillBucket(numTokens, refillRateMs) {
    this.#timer = setInterval(() => {
      this.#bucket = Math.min(this.#bucket + numTokens, this.#maxBucketSize);
    }, refillRateMs);
  }

  stopRunRefillBucket() {
    clearInterval(this.#timer);
  }
}
