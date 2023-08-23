const { productExceptSelf } = require("../index");

it.each([
  {
    nums: [1, 2, 3, 4],
    expectedOutput: [24, 12, 8, 6],
  },
  {
    nums: [-1, 1, 0, -3, 3],
    expectedOutput: [-0, 0, 9, -0, 0],
  },
  {
    nums: [1, -1],
    expectedOutput: [-1, 1],
  },
])(
  "should return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i]",
  ({ nums, expectedOutput }) => {
    const productNumsExceptSelf = productExceptSelf(nums);
    expect(productNumsExceptSelf).toEqual(expectedOutput);
  }
);
