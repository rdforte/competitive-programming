const { findTargetSumWays } = require("../index");

it("should find all expressions which evaluate to target", () => {
  const input = [1, 1, 1, 1, 1];
  const target = 3;
  const expectedOutput = 5;

  const numExpressions = findTargetSumWays(input, target);
  expect(numExpressions).toEqual(expectedOutput);
});
