const { createRateLimiter } = require("../index");

test("it should return false when requests added are below threshold", () => {
  jest.useFakeTimers().setSystemTime(new Date("2011-10-05T14:48:00.000Z"));

  const limit = 5;
  const durationMs = 1000;

  const rateLimiter = createRateLimiter(limit, durationMs);
  rateLimiter.addRequest();
  rateLimiter.addRequest();
  rateLimiter.addRequest();
  rateLimiter.addRequest();
  rateLimiter.addRequest();

  expect(rateLimiter.limitExceeded()).toBeFalsy();
});

test("it should return true when requests added are above threshold", () => {
  jest.useFakeTimers().setSystemTime(new Date("2011-10-05T14:49:00.000Z"));

  const limit = 5;
  const durationMs = 1000;

  const rateLimiter = createRateLimiter(limit, durationMs);
  rateLimiter.addRequest();
  rateLimiter.addRequest();
  rateLimiter.addRequest();
  rateLimiter.addRequest();
  rateLimiter.addRequest();
  rateLimiter.addRequest();

  expect(rateLimiter.limitExceeded()).toBeTruthy();
});

test("it should return false when requests added are above threshold but falls into new window", () => {
  jest.useFakeTimers().setSystemTime(new Date("2011-10-05T14:50:00.000Z"));

  const limit = 5;
  const durationMs = 1000;

  const rateLimiter = createRateLimiter(limit, durationMs);
  rateLimiter.addRequest();
  jest.useFakeTimers().setSystemTime(new Date("2011-10-05T14:50:00.100Z"));
  rateLimiter.addRequest();
  jest.useFakeTimers().setSystemTime(new Date("2011-10-05T14:50:00.200Z"));
  rateLimiter.addRequest();
  jest.useFakeTimers().setSystemTime(new Date("2011-10-05T14:50:00.300Z"));
  rateLimiter.addRequest();
  jest.useFakeTimers().setSystemTime(new Date("2011-10-05T14:50:00.400Z"));
  rateLimiter.addRequest();
  jest.useFakeTimers().setSystemTime(new Date("2011-10-05T14:50:01.000Z"));
  rateLimiter.addRequest();

  expect(rateLimiter.limitExceeded()).toBeFalsy();
});
