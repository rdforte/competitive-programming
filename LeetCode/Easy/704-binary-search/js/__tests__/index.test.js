const { search } = require("../index");

it.each([
  {
    nums: [-1, 0, 3, 5, 9, 12],
    target: 9,
    expectOutput: 4,
  },
  {
    nums: [-1, 0, 3, 5, 9, 12],
    target: 2,
    expectOutput: -1,
  },
  {
    nums: [5],
    target: 5,
    expectOutput: 0,
  },
])("should find the target", ({ nums, target, expectOutput }) => {
  const output = search(nums, target);
  expect(output).toEqual(expectOutput);
});
