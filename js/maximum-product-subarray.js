/**
 * @param {number[]} nums
 * @return {number}
 */
const maxProduct = (nums) => {
  if (nums.length === 0) {
    return 0;
  }

  // assume current number is the maximum
  let result = nums[0];
  let currentMax = result;
  let currentMin = result;

  // skip ahead one
  for (let i = 1; i < nums.length; i++) {
    const tempMax = currentMax;
    const currentNumber = nums[i];

    currentMax = Math.max(
      currentNumber,
      currentNumber * currentMax,
      currentNumber * currentMin
    );
    currentMin = Math.min(
      currentNumber,
      currentNumber * currentMin,
      currentNumber * tempMax
    );

    result = Math.max(currentMax, result);
  }

  return result;
};

console.log(maxProduct([2, 3, -2, 4]), 6);
console.log(maxProduct([2, 3, 4, -2, 4]), 24);
console.log(maxProduct([-2, 0, -1]), 0);
console.log(maxProduct([-2]), -2);
console.log(maxProduct([-3, -1, -1]), 3);
console.log(maxProduct([0, 2]), 2);
console.log(maxProduct([-2, 0]), 0);
console.log(maxProduct([-4, -3, -2]), 12);
console.log(maxProduct([-2, 3, -4]), 24);
console.log(maxProduct([2, -5, -2, -4, 3]), 24);
