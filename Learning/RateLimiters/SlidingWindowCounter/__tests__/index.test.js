const { createRateLimiter } = require("../index");

describe("RateLimiter", () => {
  it("should return true when exceeded rate limit", () => {
    jest.useFakeTimers().setSystemTime(new Date("2011-10-05T00:00:00.000Z"));
    const maxRequests = 5;
    const intervalMs = 1000;

    const rateLimiter = createRateLimiter(maxRequests, intervalMs);

    for (let i = 0; i < maxRequests; i++) {
      rateLimiter.addRequest();
    }

    expect(rateLimiter.exceededLimit()).toBeTruthy();
  });

  it("should return true when interval time passes and still exceed request", () => {
    jest.useFakeTimers().setSystemTime(new Date("2011-10-05T00:00:00.000Z"));
    const maxRequests = 5;
    const intervalMs = 1000;

    const rateLimiter = createRateLimiter(maxRequests, intervalMs);

    for (let i = 0; i < maxRequests; i++) {
      rateLimiter.addRequest();
    }

    jest.setSystemTime(new Date("2011-10-05T00:00:01.000Z"));

    expect(rateLimiter.exceededLimit()).toBeFalsy();
  });

  it("should return false when interval time passes and not exceeded limit", () => {
    jest.useFakeTimers().setSystemTime(new Date("2011-10-05T00:00:00.000Z"));
    const maxRequests = 5;
    const intervalMs = 1000;

    const rateLimiter = createRateLimiter(maxRequests, intervalMs);

    rateLimiter.addRequest();
    rateLimiter.addRequest();

    jest.setSystemTime(new Date("2011-10-05T00:00:01.000Z"));

    rateLimiter.addRequest();
    rateLimiter.addRequest();

    expect(rateLimiter.exceededLimit()).toBeFalsy();
  });

  it("should return false when exceeded limit and then interval time passes ", () => {
    jest.useFakeTimers().setSystemTime(new Date("2011-10-05T00:00:00.000Z"));
    const maxRequests = 5;
    const intervalMs = 1000;

    const rateLimiter = createRateLimiter(maxRequests, intervalMs);

    rateLimiter.addRequest();
    rateLimiter.addRequest();
    rateLimiter.addRequest();

    jest.setSystemTime(new Date("2011-10-05T00:00:01.000Z"));

    rateLimiter.addRequest();
    rateLimiter.addRequest();
    rateLimiter.addRequest();

    expect(rateLimiter.exceededLimit()).toBeTruthy();

    jest.setSystemTime(new Date("2011-10-05T00:00:02.000Z"));

    expect(rateLimiter.exceededLimit()).toBeFalsy();
  });

  it("should return false when exceeded limit across 2 windows and then interval time passes ", () => {
    jest.useFakeTimers().setSystemTime(new Date("2011-10-05T00:00:00.000Z"));
    const maxRequests = 5;
    const intervalMs = 1000;

    const rateLimiter = createRateLimiter(maxRequests, intervalMs);

    rateLimiter.addRequest();
    rateLimiter.addRequest();
    rateLimiter.addRequest();
    rateLimiter.addRequest();
    rateLimiter.addRequest();

    expect(rateLimiter.exceededLimit()).toBeTruthy();

    jest.setSystemTime(new Date("2011-10-05T00:00:03.000Z"));

    expect(rateLimiter.exceededLimit()).toBeFalsy();
  });
});
