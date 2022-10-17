using System;
namespace leetcode.Leetcode
{
    public class GiftingGroups
    {
        public static int countGroups(List<string> related)
        {
            int output = 0;

            if (related.Count == 0)
            {
                return output;
            }

            // create matrix
            List<int[]> relationshipMatrix = new List<int[]>();
            for (int im = 0; im < related.Count; im++)
            {
                relationshipMatrix.Add(
                    Array.ConvertAll(related[im].ToArray(), c => (int)Char.GetNumericValue(c))
                );
            }

            // iterate matrix
            int x = 0;
            int y = 0;

            for (x = 0; x < relationshipMatrix.Count;)
            {
                var currentRow = relationshipMatrix[x];
                var currentCell = relationshipMatrix[x][y];

                // iterate over row
                for (y = 0; y < currentRow.Length;)
                {
                    int matches = 0;

                    // validate top
                    if (y != 0)
                    {
                        // get one above
                        matches += relationshipMatrix[x][y-1] == 1 ? 1 : 0;
                    }

                    // validate left
                    if (x != 0)
                    {
                        matches += relationshipMatrix[x - 1][y] == 1 ? 1 : 0;
                    }

                    // validate bottom
                    if (y < relationshipMatrix.Count - 1)
                    {
                        matches += relationshipMatrix[x][y + 1] == 1 ? 1 : 0;
                    }

                    // validate right
                    if (x < currentRow.Length - 1)
                    {
                        matches += relationshipMatrix[x + 1][y] == 1 ? 1 : 0;
                    }

                    if (matches > 2)
                    {
                        output++;
                    }

                    y++;
                }

                // advance to next row
                if (x < currentRow.Length - 1)
                {
                    x = -1;
                    y = 0;
                }
                x++;
            }

            return output;
        }
    }
}

