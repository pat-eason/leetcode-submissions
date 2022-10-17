/**
 * You are given two non-empty linked lists representing
 * two non-negative integers. The digits are stored in
 * reverse order, and each of their nodes contains a single
 * digit. Add the two numbers and return the sum as a linked list.
 *
 * You may assume the two numbers do not contain any leading zero,
 * except the number 0 itself.
 *
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
const addTwoNumbers = (l1, l2) => {
  const reversedListA = l1.reverse();
  const reversedListB = l2.reverse();

  const numberA = Number.parseInt(reversedListA.join(""));
  const numberB = Number.parseInt(reversedListB.join(""));

  const total = numberA + numberB;

  return String(total)
    .split("")
    .map((x) => Number(x))
    .reverse();
};

console.log(addTwoNumbers([2, 4, 3], [5, 6, 4]), [7, 0, 8]);
console.log(addTwoNumbers([0], [0]), [0]);
console.log(
  addTwoNumbers([9, 9, 9, 9, 9, 9, 9], [9, 9, 9, 9]),
  [8, 9, 9, 9, 0, 0, 0, 1]
);
