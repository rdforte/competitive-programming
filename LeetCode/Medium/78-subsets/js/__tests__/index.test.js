const { subsets } = require("../index");

it.each([
  {
    nums: [1, 2, 3],
    expectedOutput: [[], [1], [2], [1, 2], [3], [1, 3], [2, 3], [1, 2, 3]],
  },
  {
    nums: [0],
    expectedOutput: [[], [0]],
  },
])("should calculate all subsets", ({ nums, expectedOutput }) => {
  const output = subsets(nums);
  expect(output).toHaveLength(expectedOutput.length);
  expect(output).toEqual(expect.arrayContaining(expectedOutput));
});
