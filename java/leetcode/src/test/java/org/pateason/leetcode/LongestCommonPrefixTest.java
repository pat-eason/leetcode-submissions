package org.pateason.leetcode;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

class LongestCommonPrefixTest {

  @Test
  void longestCommonPrefix_case1() {
    String[] input = new String[] {"flower","flow","flight"};
    String result = LongestCommonPrefix.longestCommonPrefix(input);
    Assertions.assertEquals("fl", result);
  }

  @Test
  void longestCommonPrefix_case2() {
    String[] input = new String[] {"dog","racecar","car"};
    String result = LongestCommonPrefix.longestCommonPrefix(input);
    Assertions.assertEquals("", result);
  }

  @Test
  void longestCommonPrefix_case3() {
    String[] input = new String[] {
        "testcase",
        "testclass",
        "testinterface",
        "testconstants",
    };
    String result = LongestCommonPrefix.longestCommonPrefix(input);
    Assertions.assertEquals("test", result);
  }

  @Test
  void longestCommonPrefix_case4() {
    String[] input = new String[] {};
    String result = LongestCommonPrefix.longestCommonPrefix(input);
    Assertions.assertEquals("", result);
  }

  @Test
  void longestCommonPrefix_case5() {
    String[] input = new String[] {""};
    String result = LongestCommonPrefix.longestCommonPrefix(input);
    Assertions.assertEquals("", result);
  }

  @Test
  void longestCommonPrefix_case6() {
    String[] input = new String[] {"ab", "a"};
    String result = LongestCommonPrefix.longestCommonPrefix(input);
    Assertions.assertEquals("a", result);
  }

  @Test
  void longestCommonPrefix_case7() {
    String[] input = new String[] {"flower", "flower", "flower", "flower"};
    String result = LongestCommonPrefix.longestCommonPrefix(input);
    Assertions.assertEquals("flower", result);
  }
}