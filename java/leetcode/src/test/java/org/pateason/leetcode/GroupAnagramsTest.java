package org.pateason.leetcode;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.List;

class GroupAnagramsTest {

  @Test
  void groupAnagrams_default_case() {
    String[] testCase = new String[] { "eat","tea","tan","ate","nat","bat" };
    List<List<String>> expectedResult = new ArrayList<>();
    expectedResult.add(new ArrayList<>(List.of("eat", "tea", "ate")));
    expectedResult.add(new ArrayList<>(List.of("bat")));
    expectedResult.add(new ArrayList<>(List.of("tan", "nat")));

    List<List<String>> result = GroupAnagrams.groupAnagrams(testCase);
    Assertions.assertEquals(expectedResult, result);
  }

  @Test
  void groupAnagrams_empty_strings() {
    String[] testCase = new String[] { "", "", "" };
    List<List<String>> expectedResult = new ArrayList<>();
    expectedResult.add(new ArrayList<>(List.of("", "", "")));

    List<List<String>> result = GroupAnagrams.groupAnagrams(testCase);
    Assertions.assertEquals(expectedResult, result);
  }

//  @Test
//  void isAnagram() {
//  }
}