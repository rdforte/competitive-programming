const { containsNearbyDuplicate } = require("../index.js");

it.each([
  {
    nums: [1, 2, 3, 1],
    k: 3,
    expectedOutput: true,
  },
  {
    nums: [1, 0, 1, 1],
    k: 1,
    expectedOutput: true,
  },
  {
    nums: [1, 2, 3, 1, 2, 3],
    k: 2,
    expectedOutput: false,
  },
])(
  "should return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k",
  ({ nums, k, expectedOutput }) => {
    const output = containsNearbyDuplicate(nums, k);
    expect(output).toEqual(expectedOutput);
  }
);
