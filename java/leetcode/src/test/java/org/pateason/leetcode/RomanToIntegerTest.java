package org.pateason.leetcode;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

class RomanToIntegerTest {

    @Test
    void romanToInt_III() {
        String testCase = "III";
        int expectedResult = 3;
        int result = RomanToInteger.romanToInt(testCase);
        Assertions.assertEquals(result, expectedResult);
    }

    @Test
    void romanToInt_LVIII() {
        String testCase = "LVIII";
        int expectedResult = 58;
        int result = RomanToInteger.romanToInt(testCase);
        Assertions.assertEquals(result, expectedResult);
    }

    @Test
    void romanToInt_MCMXCIV() {
        String testCase = "MCMXCIV";
        int expectedResult = 1994;
        int result = RomanToInteger.romanToInt(testCase);
        Assertions.assertEquals(result, expectedResult);
    }
}