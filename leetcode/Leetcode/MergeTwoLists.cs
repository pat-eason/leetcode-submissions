using System;
namespace leetcode.Leetcode;

public class MergeTwoLists
{
    public static ListNode Execute(ListNode list1, ListNode list2)
    {

        ListNode sortedTemp = new ListNode(0);
        ListNode current = sortedTemp;

        while (list1 != null && list2 != null)
        {
            if (list1.val < list2.val)
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

