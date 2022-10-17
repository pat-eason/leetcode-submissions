using System.Linq;
using leetcode.Leetcode;

internal class Program
{
    private static void Main(string[] args)
    {
        // stage variables
        // true
        ListNode test1 = new ListNode(
            1,
            new ListNode(
                2,
                new ListNode(
                    2,
                    new ListNode(1)
                )
            )
        );

        ListNode test2 = new ListNode(
            1,
            new ListNode(2)
        );

        Console.WriteLine("STARTING");

        // execute
        var result = PalindromeLinkedList.IsPalindrome(test1);

        Console.WriteLine("FINISHED");

        Console.ReadLine();
    }
}