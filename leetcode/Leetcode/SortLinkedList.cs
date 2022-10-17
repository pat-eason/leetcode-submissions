using System;
namespace leetcode.Leetcode;

public class SortLinkedList
{
    public static ListNode SortList(ListNode head)
    {
        if (head == null || head.next == null)
        {
            return head;
        }

        ListNode temp = head;
        ListNode slow = head;
        ListNode? fast = head;

        /**
         * temp will be the tail on the first half
         * slow will end on the middle of the list (+1)
         * fast will end on the tail of the list (+2)
         */
        while (fast != null && fast.next != null)
        {
            temp = slow;
            slow = slow.next;
            fast = fast.next.next;
        }

        temp.next = null;

        ListNode leftHalf = SortList(head);
        ListNode rightHalf = SortList(slow);

        return MergeList(leftHalf, rightHalf);
    }

    private static ListNode MergeList(ListNode? list1, ListNode? list2)
    {
        ListNode sortedTemp = new ListNode(0);
        ListNode current = sortedTemp;

        while (list1 != null && list2 != null)
        {
            if(list1.val < list2.val)
            {
                current.next = list1;
                list1 = list1.next;
            }
            else
            {
                current.next = list2;
                list2 = list2.next;
            }
            current = current.next;
        }

        if (list1 == null)
        {
            current.next = list2;
        }

        if (list2 == null)
        {
            current.next = list1;
        }

        return sortedTemp.next;
    }
}

