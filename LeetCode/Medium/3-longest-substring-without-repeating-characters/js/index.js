module.exports.lengthOfLongestSubstring = function (str) {
  const mp = new Map();
  let longest = 0;

  for (let l = 0, r = 0; r < str.length; r++) {
    if (mp.has(str[r])) {
      const prevRight = mp.get(str[r]);
      let left = l;
      while (left <= prevRight) {
        mp.delete(str[left]);
        left++;
      }
      l = left;
    }

    mp.set(str[l], l);

    mp.set(str[r], r);

    longest = Math.max(longest, r + 1 - l);
  }

  return longest;
};
