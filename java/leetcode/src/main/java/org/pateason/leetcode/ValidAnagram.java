package org.pateason.leetcode;

import java.util.Arrays;

public class ValidAnagram {
  public static boolean isAnagram(String s, String t) {
    char[] chars1 = s.toCharArray();
    char[] chars2 = t.toCharArray();
    Arrays.sort(chars1);
    Arrays.sort(chars2);
    return new String(chars1).equals(new String(chars2));
  }
}
