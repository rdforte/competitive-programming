const { createLRUCache } = require("../index");

it("should enable adding and removing from cache within capacity", () => {
  const capacity = 2;
  const cache = createLRUCache(capacity);
  cache.put(1, 1);
  cache.put(2, 2);
  expect(cache.get(1)).toEqual(1);
  cache.put(3, 3);
  expect(cache.get(2)).toEqual(-1);
  cache.put(4, 4);
  expect(cache.get(1)).toEqual(-1);
  expect(cache.get(3)).toEqual(3);
  expect(cache.get(4)).toEqual(4);
});
