module.exports.maxProfit = function (prices) {
  let maxProfit = 0;
  for (let l = 0, h = 0; h < prices.length; h++) {
    maxProfit = Math.max(maxProfit, prices[h] - prices[l]);

    if (prices[h] < prices[l]) {
      l = h;
    }
  }
  return maxProfit;
};
