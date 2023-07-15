const { twoSum } = require("../index.js");

it.each([
  {
    numbers: [5, 25, 75],
    target: 100,
    expectedOutput: [2, 3],
  },
  {
    numbers: [-1, 0, 1, 2, 7, 11, 15],
    target: 9,
    expectedOutput: [4, 5],
  },
  {
    numbers: [2, 7, 11, 15],
    target: 9,
    expectedOutput: [1, 2],
  },
  {
    numbers: [2, 3, 4],
    target: 6,
    expectedOutput: [1, 3],
  },
  {
    numbers: [-1, 0],
    target: -1,
    expectedOutput: [1, 2],
  },
])(
  "should retrieve the index's used to calculate the target",
  ({ numbers, target, expectedOutput }) => {
    const indexes = twoSum(numbers, target);
    expect(indexes).toEqual(expectedOutput);
  }
);
