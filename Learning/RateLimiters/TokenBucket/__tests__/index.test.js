const { createRateLimiter } = require("../index");

it("should return true when exceeded rate limit", () => {
  const bucketSize = 3;
  const refillNumTokensPerSecond = 2;
  const rateLimiter = createRateLimiter(bucketSize, refillNumTokensPerSecond);

  jest.useFakeTimers().setSystemTime(new Date("2011-10-05T14:48:00.000Z"));

  exhaustBucket(rateLimiter, bucketSize);

  expect(rateLimiter.exceededLimit()).toBeTruthy();
});

it("should return false when bucket is drained and then bucket refills", () => {
  jest.useFakeTimers();
  jest.setSystemTime(new Date("2011-10-05T14:48:00.000Z"));

  const bucketSize = 3;
  const refillNumTokensPerSecond = 2;
  const rateLimiter = createRateLimiter(bucketSize, refillNumTokensPerSecond);

  exhaustBucket(rateLimiter, bucketSize);

  jest.setSystemTime(new Date("2011-10-05T14:48:01.000Z"));

  expect(rateLimiter.exceededLimit()).toBeFalsy();
});

it("should only refill to the max bucket size", () => {
  jest.useFakeTimers().setSystemTime(new Date("2011-10-05T14:48:00.000Z"));

  const bucketSize = 3;
  const refillNumTokensPerSecond = 2;
  const rateLimiter = createRateLimiter(bucketSize, refillNumTokensPerSecond);

  exhaustBucket(rateLimiter, bucketSize);

  const advanceTime = 1000 * bucketSize;
  jest.advanceTimersByTime(advanceTime);

  exhaustBucket(rateLimiter, bucketSize);

  expect(rateLimiter.exceededLimit()).toBeTruthy();
});

const exhaustBucket = (rateLimiter, bucketSize) => {
  for (let i = 0; i < bucketSize; i++) {
    rateLimiter.take();
  }
};
