using System;
namespace leetcode.Leetcode
{
    public class SquaresOfASortedArray
    {
        public static int[] SortedSquares(int[] nums)
        {
            int[] temp = new int[nums.Length];

            for (int i = 0; i < nums.Length; i++)
            {
                temp[i] = nums[i] * nums[i];
            }

            HeapSort(temp, temp.Length);

            return temp;
        }

        private static void HeapSort(int[] arr, int n)
        {
            for (int i = n / 2 - 1; i >= 0; i--)
            {
                Heapify(arr, n, i);
            }
                
            for (int i = n - 1; i >= 0; i--)
            {
                int temp = arr[0];
                arr[0] = arr[i];
                arr[i] = temp;
                Heapify(arr, i, 0);
            }
        }

        private static void Heapify(int[] arr, int n, int i)
        {
            int largest = i;
            int left = 2 * i + 1;
            int right = 2 * i + 2;
            if (left < n && arr[left] > arr[largest])
            {
                largest = left;
            }
                
            if (right < n && arr[right] > arr[largest])
            {
                largest = right;
            }
                
            if (largest != i)
            {
                int swap = arr[i];
                arr[i] = arr[largest];
                arr[largest] = swap;
                Heapify(arr, n, largest);
            }
        }
    }
}

