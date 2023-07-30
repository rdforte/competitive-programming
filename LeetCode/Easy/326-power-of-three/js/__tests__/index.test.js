const { isPowerOfThree } = require("../index");

it.each([
  {
    n: 27,
    expectedOutput: true,
  },
  {
    n: 0,
    expectedOutput: false,
  },
  {
    n: -1,
    expectedOutput: false,
  },
  {
    n: 1,
    expectedOutput: true,
  },
  {
    n: -3,
    expectedOutput: false,
  },
  {
    n: -9,
    expectedOutput: false,
  },
])("should determine if $n is power of 3", ({ n, expectedOutput }) => {
  const output = isPowerOfThree(n);
  expect(output).toEqual(expectedOutput);
});
