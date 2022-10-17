using System;
namespace leetcode.Leetcode
{
    public class OptimizingBoxWeights
    {
        public static List<int> minimalHeaviestSetA(List<int> arr)
        {
            // fallback if there cannot be a subset. List is empty or only has one item
            if (arr.Count < 2)
            {
                return arr;
            }

            // sort the array in descending order
            arr.Sort((a, b) => a.CompareTo(b));

            // start a stack, track the remaining weight, and create an output list
            Stack<int> stack = new Stack<int>();
            int remainingWeight = 0;
            List<int> output = new List<int>();

            // fill the stack wth the current array and assign the total weight
            for (int i = 0; i < arr.Count; i++)
            {
                remainingWeight += arr[i];
                stack.Push(arr[i]);
            }

            // iterate over the stack
            while (stack.Count > 0)
            {
                int current = stack.Pop();

                // get the sum to insert
                output.Insert(0, current);

                // using linq because it's convenient
                int currentTotalSum = output.Sum();

                remainingWeight -= current;

                // if the curent is less than the remaining weight, with the current total removed, then break
                if (!(current < (remainingWeight - currentTotalSum)))
                {
                    break;
                }
            }
            return output;
        }
    }
}

