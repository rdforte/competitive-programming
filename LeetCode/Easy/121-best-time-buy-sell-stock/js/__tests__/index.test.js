const { maxProfit } = require("../index");

it.each([
  {
    prices: [7, 1, 5, 3, 6, 4],
    expectedOutput: 5,
  },
  {
    prices: [7, 6, 4, 3, 1],
    expectedOutput: 0,
  },
])("should calculate the max profit", ({ prices, expectedOutput }) => {
  const profit = maxProfit(prices);
  expect(profit).toEqual(expectedOutput);
});
