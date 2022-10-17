using System;
namespace leetcode.Leetcode
{
    public class BinaryLinkedListToDecimal
    {
        public static int Execute(ListNode head)
        {
            if (head == null)
            {
                return 0;
            }

            // start the decimal
            int output = 0;

            // traverse the tree
            ListNode current = head;
            while (current != null)
            {
                // multiply output by two and add node value
                output = (output * 2) + current.val;
                current = current.next;
            }

            return output;
        }
    }
}

