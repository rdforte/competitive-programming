module.exports.findWinnerOfElection = (votes) => {
  const voteCount = new Map();

  let winner = votes[0];

  for (let v of votes) {
    if (voteCount.has(v)) {
      voteCount.set(v, voteCount.get(v) + 1);
    } else {
      voteCount.set(v, 1);
    }

    const candidateVotes = voteCount.get(v);
    const winnerVotes = voteCount.get(winner);

    const candidateHasMoreVotes = candidateVotes > winnerVotes;
    const candidateAndWinnerHaveEqualVotes = candidateVotes === winnerVotes;

    if (candidateHasMoreVotes) {
      winner = v;
    }
    if (candidateAndWinnerHaveEqualVotes) {
      winner = [winner, v].sort()[0];
    }
  }

  return [winner, voteCount.get(winner)];
};
