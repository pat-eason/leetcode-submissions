package org.pateason.leetcode;

/**
 * @see <a href="https://leetcode.com/problems/zigzag-conversion/">Leetcode</a>
 *
 * The string "PAYPALISHIRING" is written in a zigzag pattern on a
 * given number of rows like this: (you may want to display this
 * pattern in a fixed font for better legibility)
 *
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 * And then read line by line: "PAHNAPLSIIGYIR"
 *
 * Write the code that will take a string and make this
 * conversion given a number of rows:
 */
public class ZigzagConversion {
  public static String convert(String s, int numRows) {
    if (numRows == 1 || numRows > s.length()) {
      return s;
    }

    String[] outputArray = new String[numRows];

    int n = 0;
    boolean reverse = false;
    for (int i = 0; i < s.length(); i++) {
      outputArray[n] = outputArray[n] == null
        ? "" + s.charAt(i)
        : outputArray[n] + s.charAt(i);

      if (reverse) {
        n--;
      } else {
        n++;
      }

      if (n == numRows - 1) {
        reverse = true;
      }

      if (n == 0) {
        reverse = false;
      }
    }

    return String.join("", outputArray);
  }
}
