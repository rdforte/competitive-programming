module.exports.isPalindrome = (str) => {
  const formattedStr = str.toLowerCase().replace(/[^a-zA-Z\d]/g, "");

  const strLen = formattedStr.length;
  if (strLen <= 1) return true;
  const middleIndex = formattedStr.length / 2;

  for (let i = 0; i < middleIndex; i++) {
    const left = formattedStr[i];
    const right = formattedStr[strLen - 1 - i];

    if (left !== right) return false;
  }
  return true;
};
