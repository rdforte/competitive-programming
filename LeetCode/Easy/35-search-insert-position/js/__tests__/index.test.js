const { searchInsert } = require("../index");

it.each([
  {
    nums: [1, 3, 5, 6],
    target: 5,
    expectedOutput: 2,
  },
  {
    nums: [1, 3, 5, 6],
    target: 2,
    expectedOutput: 1,
  },
  {
    nums: [1, 3, 5, 6],
    target: 7,
    expectedOutput: 4,
  },
])(
  "should return the index if the target is found",
  ({ nums, target, expectedOutput }) => {
    const output = searchInsert(nums, target);
    expect(output).toEqual(expectedOutput);
  }
);
