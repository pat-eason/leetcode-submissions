/**
 * @param {number[]} nums
 * @return {number}
 */
const maximumProduct = (nums) => {
  const sorter = (a, b) => a - b;
  nums.sort(sorter);
  let pro1 = 1;
  let pro2 = 1;
  let len = nums.length - 1;
  for (let i = len; i > len - 3; i--) {
    pro1 = pro1 * nums[i];
  }
  pro2 = nums[0] * nums[1] * nums[len];
  return Math.max(pro1, pro2);
};

console.log(maximumProduct([-100, -98, -1, 2, 3, 4]), 39200);
console.log(maximumProduct([-100, -2, -3, 1]), 300);
console.log(maximumProduct([-1, -2, 1, 2, 3]), 6);
console.log(maximumProduct([-8, -7, -2, 10, 20]), 1120);
