/**
 * @param {number[]} nums
 * @return {number}
 */
const maxSubArray = (nums) => {
  if (nums.length === 0) {
    return 0;
  }

  // assume current number is the maximum
  let result = nums[0];
  let currentMax = 0;

  // skip ahead one
  for (let i = 0; i < nums.length; i++) {
    const currentNumber = nums[i];
    currentMax = Math.max(currentNumber, currentNumber + currentMax);
    result = Math.max(result, currentMax);
  }

  return result;
};

console.log(maxSubArray([-2, 1, -3, 4, -1, 2, 1, -5, 4]), 6);
console.log(maxSubArray([1]), 1);
console.log(maxSubArray([5, 4, -1, 7, 8]), 23);
