package org.pateason.leetcode;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

class FindPlayersWithZeroOrOneLossesTest {

  @Test
  void findWinners_case1() {
    int[][] input = new int[][]{
        {1, 3},
        {2, 3},
        {3, 6},
        {5, 6},
        {5, 7},
        {4, 5},
        {4, 8},
        {4, 9},
        {10, 4},
        {10, 9}
    };
    List<List<Integer>> expected = new ArrayList<>(
        Arrays.asList(
            Arrays.asList(1, 2, 10),
            Arrays.asList(4, 5, 7, 8)
        )
    );
    List<List<Integer>> result = FindPlayersWithZeroOrOneLosses.findWinners(input);
    Assertions.assertEquals(expected, result);
  }

  @Test
  void findWinners_case2() {
    int[][] input = new int[][]{
        {2, 3}, {1, 3}, {5, 4}, {6, 4}
    };
    List<List<Integer>> expected = new ArrayList<>(
        Arrays.asList(
            Arrays.asList(1, 2, 5, 6),
            List.of()
        )
    );
    List<List<Integer>> result = FindPlayersWithZeroOrOneLosses.findWinners(input);
    Assertions.assertEquals(expected, result);
  }
}