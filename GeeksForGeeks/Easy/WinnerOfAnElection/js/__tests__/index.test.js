const { findWinnerOfElection } = require("../index");

it.each([
  {
    votes: [
      "john",
      "johnny",
      "jackie",
      "johnny",
      "john",
      "jackie",
      "jamie",
      "jamie",
      "john",
      "johnny",
      "jamie",
      "johnny",
      "john",
    ],
    expectedWinner: ["john", 4],
  },
])("should find the winner of the election", ({ votes, expectedWinner }) => {
  const winner = findWinnerOfElection(votes);
  expect(winner).toEqual(expectedWinner);
});
