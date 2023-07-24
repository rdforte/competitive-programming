module.exports.threeSum = (nums) => {
  const n = nums.sort((a, b) => a - b);
  const duplicates = new Set();

  const res = [];

  for (let i = 0; i < n.length - 2; i++) {
    const numI = n[i];
    for (let l = i + 1, r = n.length - 1; l < r; ) {
      const numL = n[l];
      const numR = n[r];

      const total = numI + numL + numR;

      if (total === 0) {
        if (!duplicates.has(`${numI}${numL}${numR}`)) {
          res.push([numI, numL, numR]);
        }
        duplicates.add(`${numI}${numL}${numR}`);
        l++;
        r--;
      } else if (total > 0) {
        r--;
      } else {
        l++;
      }
    }
  }

  return res;
};
