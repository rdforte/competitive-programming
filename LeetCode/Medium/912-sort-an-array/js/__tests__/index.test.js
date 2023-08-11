const { sortArray } = require("../index");

it.each([
  // {
  //   nums: [5, 2, 3, 1],
  //   expectedOutput: [1, 2, 3, 5],
  // },
  {
    nums: [5, 1, 1, 2, 0, 0],
    expectedOutput: [0, 0, 1, 1, 2, 5],
  },
])("should sort the array $nums", ({ nums, expectedOutput }) => {
  const sortedNums = sortArray(nums);
  expect(sortedNums).toEqual(expectedOutput);
});
