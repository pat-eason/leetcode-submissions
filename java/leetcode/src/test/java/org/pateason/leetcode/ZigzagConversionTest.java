package org.pateason.leetcode;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class ZigzagConversionTest {

  @Test
  void convert_PAYPALISHIRING_4_rows() {
    int numberOfRows = 4;
    String input = "PAYPALISHIRING";
    String expectedResult = "PINALSIGYAHRPI";

    String result = ZigzagConversion.convert(input, numberOfRows);
    assertEquals(expectedResult, result);
  }

  @Test
  void convert_PAYPALISHIRING_3_rows() {
    int numberOfRows = 3;
    String input = "PAYPALISHIRING";
    String expectedResult = "PAHNAPLSIIGYIR";

    String result = ZigzagConversion.convert(input, numberOfRows);
    assertEquals(expectedResult, result);
  }

  @Test
  void convert_AB_1_row() {
    int numberOfRows = 1;
    String input = "AB";
    String expectedResult = "AB";

    String result = ZigzagConversion.convert(input, numberOfRows);
    assertEquals(expectedResult, result);
  }

  @Test
  void convert_A_2_row() {
    int numberOfRows = 2;
    String input = "A";
    String expectedResult = "A";

    String result = ZigzagConversion.convert(input, numberOfRows);
    assertEquals(expectedResult, result);
  }
}