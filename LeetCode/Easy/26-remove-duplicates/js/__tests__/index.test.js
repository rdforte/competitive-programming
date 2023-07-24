const { removeDuplicates } = require("../index");

it.each([
  {
    nums: [1, 1, 2],
    expectedLenOutput: 2,
    expectedNumsOutput: [1, 2],
  },
  {
    nums: [0, 0, 1, 1, 1, 2, 2, 3, 3, 4],
    expectedLenOutput: 5,
    expectedNumsOutput: [0, 1, 2, 3, 4],
  },
])(
  "should return the number of duplicates",
  ({ nums, expectedLenOutput, expectedNumsOutput }) => {
    const removedNumDuplicates = removeDuplicates(nums);

    for (let i = 0; i < expectedLenOutput; i++) {
      expect(nums[i]).toEqual(expectedNumsOutput[i]);
    }
  }
);
