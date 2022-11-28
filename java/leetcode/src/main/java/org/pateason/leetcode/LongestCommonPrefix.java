package org.pateason.leetcode;

/**
 * @see <a href="https://leetcode.com/problems/longest-common-prefix/">Leetcode</a>
 *
 * Write a function to find the longest common prefix string
 * amongst an array of strings.
 *
 * If there is no common prefix, return an empty string "".
 */
public class LongestCommonPrefix {
  public static String longestCommonPrefix(String[] strs) {
    if (strs.length == 0) {
      return "";
    }

    StringBuilder output = new StringBuilder();

    // create matrix
    char[][] charMatrix = new char[strs.length][];
    for (int i = 0; i < strs.length; i++) {
      charMatrix[i] = strs[i].toCharArray();
    }

    // iterate and compare
    char[] firstWord = charMatrix[0];
    for (int i = 0; i < firstWord.length; i++) {
      boolean allMatch = true;
      char comparisonChar = firstWord[i];

      for (int j = 1; j < charMatrix.length; j++) {
        if (charMatrix[j].length == i) {
          allMatch = false;
          break;
        }

        char targetChar = charMatrix[j][i];
        if (targetChar != comparisonChar) {
          allMatch = false;
          break;
        }
      }

      if (allMatch) {
        output.append(comparisonChar);
      }
      if (!allMatch) {
        break;
      }
    }

    return output.toString();
  }
}
