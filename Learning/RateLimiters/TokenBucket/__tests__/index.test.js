const { createRateLimiter } = require("../index");

it("should return true when exceeded rate limit", () => {
  const bucketSize = 3;

  const rateLimiter = createRateLimiter(bucketSize);
  const numTokens = 2;
  const refillRateMs = 1000;
  rateLimiter.startRunRefillBucket(numTokens, refillRateMs);

  jest.useFakeTimers().setSystemTime(new Date("2011-10-05T14:48:00.000Z"));

  exhaustBucket(rateLimiter, bucketSize);

  expect(rateLimiter.exceededLimit()).toBeTruthy();
});

it("should return false when bucket is drained and then bucket refills", () => {
  jest.useFakeTimers().setSystemTime(new Date("2011-10-05T14:48:00.000Z"));

  const bucketSize = 3;
  const rateLimiter = createRateLimiter(bucketSize);
  const numTokens = 2;
  const refillRateMs = 1000;
  rateLimiter.startRunRefillBucket(numTokens, refillRateMs);

  exhaustBucket(rateLimiter, bucketSize);

  jest.advanceTimersByTime(refillRateMs);

  expect(rateLimiter.exceededLimit()).toBeFalsy();

  rateLimiter.stopRunRefillBucket();
});

it("should only refill to the max bucket size", () => {
  jest.useFakeTimers().setSystemTime(new Date("2011-10-05T14:48:00.000Z"));

  const bucketSize = 3;
  const rateLimiter = createRateLimiter(bucketSize);
  const numTokens = 2;
  const refillRateMs = 1000;
  rateLimiter.startRunRefillBucket(numTokens, refillRateMs);

  exhaustBucket(rateLimiter, bucketSize);

  const advanceTime = refillRateMs * bucketSize;
  jest.advanceTimersByTime(advanceTime);

  exhaustBucket(rateLimiter, bucketSize);

  expect(rateLimiter.exceededLimit()).toBeTruthy();

  rateLimiter.stopRunRefillBucket();
});

const exhaustBucket = (rateLimiter, bucketSize) => {
  for (let i = 0; i < bucketSize; i++) {
    rateLimiter.addRequest();
  }
};
