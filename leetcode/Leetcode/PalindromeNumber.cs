using System;
namespace leetcode.Leetcode
{
    public class PalindromeNumber
    {
        public static bool IsPalindrome(int x)
        {
            bool output = true;

            char[] xCharArray = x.ToString().ToCharArray();

            int i = 0;
            int j = xCharArray.Length - 1;

            while (i <= j)
            {
                if (xCharArray[i] != xCharArray[j])
                {
                    output = false;
                    break;
                }
                i++;
                j--;
            }

            return output;
        }
    }
}

