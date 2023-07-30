module.exports.isPowerOfThree = (n) => calculateIsPowerOfThree(n);

const calculateIsPowerOfThree = (n) => {
  if (n === 1) return true;
  else if (n < 1) return false;
  return calculateIsPowerOfThree(n / 3);
};
