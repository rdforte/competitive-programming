module.exports.subsets = (nums) => {
  const subsets = [];
  const res = [];
  calculateSubsetsDfsV2(nums, 0, res, subsets);
  return subsets;
};

// faster and more optimal on memory as we pass res and subsets by reference.
const calculateSubsetsDfsV2 = (nums, index, res, subsets) => {
  if (index >= nums.length) {
    subsets.push([...res]);
    return;
  }

  res.push(nums[index]);
  calculateSubsetsDfsV2(nums, index + 1, res, subsets);
  res.pop();
  calculateSubsetsDfsV2(nums, index + 1, res, subsets);
};

// less optimal but more straight forward as we just bubble each result up and add to array and then bubble up again.
const calculateSubsetsDfs = (nums, index, res) => {
  if (index >= nums.length) {
    return [res];
  }

  const subsets = [];

  const withoutNum = calculateSubsetsDfs(nums, index + 1, [...res]);
  withoutNum.forEach((w) => subsets.push(w));

  const withNum = calculateSubsetsDfs(nums, index + 1, [...res, nums[index]]);
  withNum.forEach((w) => subsets.push(w));

  return subsets;
};
