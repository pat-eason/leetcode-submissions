using System;
namespace leetcode.Leetcode
{
    public class ReverseLinkedList
    {
        public static ListNode ReverseList(ListNode head)
        {
            ListNode curr = head;
            ListNode prev = null;
            ListNode next = null;
            while (curr != null)
            {
                // hold next as temp
                next = curr.next;

                // update current next to the previous (head - 1)
                curr.next = prev;

                // update the previous to the current (head)
                prev = curr;

                // update the next to the current
                curr = next;
            }

            return prev;
        }
    }
}

