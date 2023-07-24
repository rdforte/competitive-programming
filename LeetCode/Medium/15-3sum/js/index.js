module.exports.threeSum = (nums) => {
  const seen = Array(nums.length).fill(false);
  const validComboIndexes = Array(nums.length).fill(false);
  return calculateThreeSumRecursively(nums, seen, 1, 0, [], validComboIndexes);
};

const calculateThreeSumRecursively = (
  nums,
  seen,
  level,
  total,
  perm,
  validComboIndexes
) => {
  const permutations = [];
  for (let i = 0; i < nums.length; i++) {
    if (validComboIndexes[i] || seen[i]) {
      continue;
    }

    seen[i] = true;
    if (level === 3 && nums[i] + total === 0) {
      permutations.push([...perm, nums[i]]);
      validComboIndexes[i] = true;
    }
    const calculatedPerms = calculateThreeSumRecursively(
      nums,
      seen,
      level + 1,
      total + nums[i],
      [...perm, nums[i]],
      validComboIndexes
    );
    calculatedPerms.forEach((perm) => permutations.push(perm));
    seen[i] = validComboIndexes[i];
  }

  return permutations;
};
