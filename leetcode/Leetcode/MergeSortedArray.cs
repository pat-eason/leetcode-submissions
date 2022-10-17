using System;
namespace leetcode.Leetcode
{
    public class MergeSortedArray
    {
        public static void Merge(int[] nums1, int m, int[] nums2, int n)
        {
            int[] temp = new int[m + n];

            int indexA = 0;
            int indexB = 0;
            int tempIndex = 0;

            while (indexA < m && indexB < n)
            {
                if (nums1[indexA] < nums2[indexB])
                {
                    temp[tempIndex++] = nums1[indexA++];
                }
                else
                {
                    temp[tempIndex++] = nums2[indexB++];
                }
            }

            while (indexA < m)
            {
                temp[tempIndex++] = nums1[indexA++];
            }

            while (indexB < n)
            {
                temp[tempIndex++] = nums2[indexB++];
            }

            for (int i = 0; i < nums1.Length; i++)
            {
                nums1[i] = temp[i];
            }


            Console.WriteLine("Done");
        }
    }
}

