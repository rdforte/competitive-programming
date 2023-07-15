module.exports.twoSum = (numbers, target) => {
  let lowerVal, upperVal;
  for (let left = 0, right = numbers.length - 1; left < right; ) {
    lowerVal = numbers[left];
    upperVal = numbers[right];
    const sum = lowerVal + upperVal;
    if (sum === target) {
      return [left + 1, right + 1];
    }
    if (sum > target) {
      right--;
    } else {
      left++;
    }
  }

  return [];
};
