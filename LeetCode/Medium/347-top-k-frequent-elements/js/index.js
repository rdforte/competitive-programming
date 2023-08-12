const nums = [1, 1, 1, 2, 2, 3],
  k = 2;

const topKFrequent = function (nums, k) {
  const elements = new Map(Array.from(new Set(nums)).map((num) => [num, 0]));

  for (let n of nums) {
    elements.set(n, elements.get(n) + 1);
  }

  const sorted = [];
  for (let [key, value] of elements.entries()) {
    sorted.push({ key, value });
  }

  sorted.sort((a, b) => a.value - b.value);

  return sorted.slice(-k).map(({ key }) => key);
};

console.log(topKFrequent(nums, k));
