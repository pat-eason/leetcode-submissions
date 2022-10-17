using System;
using System.Linq;
using System.Text.RegularExpressions;

namespace leetcode.Leetcode
{
    public class IsValidPalindrome
    {
        public static bool Execute(string input)
        {
            string sanitizedInput = Regex.Replace(input, "[^a-zA-Z0-9]", string.Empty).ToLower();
            string reversedSanitizedInput = string.Join("", sanitizedInput.Reverse());

            return sanitizedInput.Length == 0 || sanitizedInput == reversedSanitizedInput;
        }
    }
}

