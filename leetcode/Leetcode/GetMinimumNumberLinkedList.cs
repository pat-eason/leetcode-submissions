using System;
namespace leetcode.Leetcode
{
    public class GetMinimumIntLinkedList
    {
        public static int Execute(LinkedList<int> input)
        {
            if (input.First == null)
            {
                return 0;
            }

            LinkedListNode<int>? current = input.First;
            int output = input.First.Value;

            while (current != null)
            {
                output = Math.Min(output, current.Value);
                current = current.Next;
            }

            return output;
        }
    }
}

