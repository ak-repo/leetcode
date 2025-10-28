/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var findMaxAverage = function(nums, k) {

    let wSum = 0;
    let mSum = -Infinity;

    // initial window
    for (let i = 0; i < k; i++) {
      wSum += nums[i];
    }
    mSum = wSum;

    // slide the window
    for (let i = k; i < nums.length; i++) {
      wSum = wSum - nums[i - k] + nums[i];
      mSum = Math.max(mSum, wSum);
    }

    return mSum/k;
    
};