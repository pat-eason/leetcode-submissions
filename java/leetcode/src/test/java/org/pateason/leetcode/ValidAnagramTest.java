package org.pateason.leetcode;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

class ValidAnagramTest {

  @Test
  void isAnagram_anagram_nagaram() {
    Assertions.assertTrue(ValidAnagram.isAnagram("anagram", "nagaram"));
  }

  @Test
  void isAnagram_rat_car() {
    Assertions.assertFalse(ValidAnagram.isAnagram("rat", "car"));
  }
}