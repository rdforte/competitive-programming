module.exports.longestConsecutiveSequence = (nums) => {
  const st = new Set();

  for (let n of nums) {
    st.add(n);
  }

  let count = 0;

  for (let n of nums) {
    // by checking that there is no previous we know that we have our starting point. If there
    // is a previous then we are probably starting off somewhere in the middle of the sequence.
    // This one line brings the algorithm down form O(n^2) to O(n) time
    if (st.has(n - 1)) continue;

    let currentCount = 1;
    while (st.has(n + 1)) {
      n++;
      currentCount++;
    }
    count = Math.max(count, currentCount);
  }

  return count;
};
