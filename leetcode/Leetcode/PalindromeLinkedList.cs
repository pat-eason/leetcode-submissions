using System;
namespace leetcode.Leetcode
{
    public class PalindromeLinkedList
    {
        public static bool IsPalindrome(ListNode head)
        {
            if (head == null)
            {
                return false;
            }

            bool output = true;

            // get middle of list
            ListNode slow = head;
            ListNode fast = head;

            /**
             * advance in three speeds
             *  - slow advances one
             *  - fast advances two
             *  
             *  at end of iteration slow should be the head of the
             *  middle of the second half of the list, fast will
             *  be the tail
             */
            while (fast != null && fast.next != null)
            {
                slow = slow.next;
                fast = fast.next.next;
            }

            /**
             * invert the slow list
             * 
             * track three temp pointers: current head, previous item, next item
             * iterate until the current head is null
             * 
             * during iteration:
             *  - set the next pointer to the next value of the current head (slow + 1) as a temp
             *  - set the next item of the current pointer to the previous item (swapping slow + 1 with slow)
             *  - set the previous pointer to the current pointer
             *  - set the current pointer to the next pointer
             */
            ListNode current = slow;
            ListNode? previous = null;
            ListNode? next = null;
            while (current != null)
            {
                next = current.next;
                current.next = previous;
                previous = current;
                current = next;
            }
            slow = previous;

            /**
             * iterate until slow is exhausted
             **/
            while (slow != null)
            {
                // compare values. if false then set the output to false and stop iterating

                if (head.val != slow.val)
                {
                    output = false;
                    break;
                }

                // advance both lists
                head = head.next;
                slow = slow.next;
            }

            return output;
        }
    }
}

