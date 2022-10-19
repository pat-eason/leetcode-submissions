package org.pateason.leetcode;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

class IntegerToRomanTest {

    @Test
    void intToRoman_3() {
        int testCase = 3;
        String expectedResult = "III";
        String result = IntegerToRoman.intToRoman(testCase);
        Assertions.assertEquals(result, expectedResult);
    }

    @Test
    void intToRoman_58() {
        int testCase = 58;
        String expectedResult = "LVIII";
        String result = IntegerToRoman.intToRoman(testCase);
        Assertions.assertEquals(result, expectedResult);
    }

    @Test
    void intToRoman_1994() {
        int testCase = 1994;
        String expectedResult = "MCMXCIV";
        String result = IntegerToRoman.intToRoman(testCase);
        Assertions.assertEquals(result, expectedResult);
    }

    @Test
    void intToRoman_20() {
        int testCase = 20;
        String expectedResult = "XX";
        String result = IntegerToRoman.intToRoman(testCase);
        Assertions.assertEquals(result, expectedResult);
    }
}