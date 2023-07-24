const { catMouseGame } = require("../index");

it.each([
  {
    input: [[2, 5], [3], [0, 4, 5], [1, 4, 5], [2, 3], [0, 2, 3]],
    expectOutput: 0,
  },
  {
    input: [[1, 3], [0], [3], [0, 2]],
    expectOutput: 1,
  },
])("should tell you who won the game", ({ input, expectOutput }) => {
  const out = catMouseGame(input);
  expect(out).toEqual(expectOutput);
});
