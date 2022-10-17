using System;
namespace leetcode.Leetcode
{
    public class CheckIfTheSentenceIsPangram
    {
        public static bool CheckIfPangram(string sentence)
        {
            if (sentence.Length < 2)
            {
                return false;
            }

            int i = 0;
            int n = sentence.Length - 1;
            List<char> pangramChars = new List<char>();
            char[] sentenceChars = sentence.ToCharArray();

            while (i <= n)
            {
                char leftChar = sentenceChars[i];
                char rightChar = sentenceChars[n];

                if (!pangramChars.Contains(leftChar))
                {
                    pangramChars.Add(leftChar);
                }
                if (!pangramChars.Contains(rightChar))
                {
                    pangramChars.Add(rightChar);
                }
                i++;
                n--;
            }

            return pangramChars.Count == 26;
        }
    }
}

