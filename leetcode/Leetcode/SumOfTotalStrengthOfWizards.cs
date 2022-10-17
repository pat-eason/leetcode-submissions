using System;
namespace leetcode.Leetcode
{
    public class SumOfTotalStrengthOfWizards
    {
        private static long _modulo = 1000000007;
        public static int TotalStrength(int[] strength)
        {
            // if there are no records, then there is nothing to calculate
            // return 0
            if (strength.Length == 0)
            {
                return 0;
            }

            long output = DynamicArray(strength);

            return (int) (output % _modulo);
        }

        private static long DynamicArray(int[] strength)
        {
            long output = 0;
            int n = strength.Length;

            // keep track of the calculated sums and products
            long[] leftSum = new long[n + 1];
            long[] leftProduct = new long[n + 1];
            long[] rightSum = new long[n + 1];
            long[] rightProduct = new long[n + 1];

            // fill the tracking arrays
            for (int i = 0; i < n; i++)
            {
                leftSum[i + 1] = (leftSum[i] + strength[i]) % _modulo;
                leftProduct[i + 1] = (leftProduct[i] + (long) (i + 1) * strength[i]) % _modulo;
            }

            for (int i = n - 1; i >= 0; i--)
            {
                rightSum[i] = (rightSum[i + 1] + strength[i]) % _modulo;
                rightProduct[i] = (rightProduct[i + 1] + (long)(n - i) * strength[i]) % _modulo;
            }

            // init stack and iterate over array
            Stack<int> stack = new Stack<int>();
            for (int rIdx = 0; rIdx <= n; rIdx++)
            {
                // while stack is not empty and the top of the stack is not less than current item
                while (
                    stack.Count > 0 && (
                        rIdx == n || strength[stack.Peek()] >= strength[rIdx]
                    )
                )
                {
                    // get the current top of the stack
                    int currentIdx = stack.Pop();
                    // get the next stack item to compare to
                    int lIdx = stack.Count == 0
                        ? 0
                        : stack.Peek() + 1;

                    long lSum = (
                        _modulo + leftProduct[currentIdx + 1]
                        - leftProduct[lIdx] - lIdx
                        * (leftSum[currentIdx + 1] - leftSum[lIdx])
                        % _modulo
                    ) % _modulo;

                    long rSum = (
                        _modulo + rightProduct[currentIdx + 1]
                        - rightProduct[rIdx] - (n - rIdx)
                        * (rightSum[currentIdx + 1] - rightSum[rIdx])
                    ) % _modulo;

                    long sum = (
                        rSum * (currentIdx - lIdx + 1)
                        + lSum * (rIdx - currentIdx)
                    ) % _modulo;
                    output = (output + sum * strength[currentIdx]) % _modulo;
                }
                stack.Push(rIdx);
            }
            return output;
        }

        private static long BruteForce(int[] strength)
        {
            long output = 0;
            int n = strength.Length;

            // do iteration here
            for (int i = 0; i < n; i++)
            {
                var currentMin = strength[i];
                long currentSum = 0;

                for (int j = 0; j < n - i; j++)
                {
                    var subStrength = strength[j + i];
                    if (subStrength < currentMin)
                    {
                        currentMin = subStrength;
                    }
                    currentSum += subStrength;

                    var powerTotal = currentMin * currentSum;

                    output += powerTotal;
                }
            }
            return output;
        }
    }
}

