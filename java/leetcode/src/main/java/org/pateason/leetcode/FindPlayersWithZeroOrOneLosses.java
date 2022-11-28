package org.pateason.leetcode;

import java.util.*;
import java.util.stream.Collectors;

/**
 * @see <a href="https://leetcode.com/problems/find-players-with-zero-or-one-losses/">Leetcode</a>
 *
 * You are given an integer array matches where <code>matches[i] = [winneri, loseri]</code>
 * indicates that the player <code>winneri</code> defeated player <code>loseri</code> in a match.
 * <br/>
 * Return a list answer of size 2 where:
 * <ul>
 *   <li>answer[0] is a list of all players that have not lost any matches.</li>
 *   <li>answer[1] is a list of all players that have lost exactly one match.</li>
 * </ul>
 * The values in the two lists should be returned in increasing order.
 * <br/>
 * Note:
 * <ul>
 *  <li>You should only consider the players that have played at least one match.</li>
 *  <li>The testcases will be generated such that no two matches will have the same outcome.</li>
 * </ul>
 */
public class FindPlayersWithZeroOrOneLosses {
  public static List<List<Integer>> findWinners(int[][] matches) {
    List<List<Integer>> output = new ArrayList<>();

    Set<Integer> players = new TreeSet<>();
    HashMap<Integer, Integer> winCounts = new HashMap<>();
    HashMap<Integer, Integer> lossCounts = new HashMap<>();

    for (int[] match : matches) {
      int winner = match[0];
      int loser = match[1];

      players.add(winner);
      players.add(loser);

      winCounts.put(winner, winCounts.getOrDefault(winner, 0) + 1);
      lossCounts.put(loser, lossCounts.getOrDefault(loser, 0) + 1);
    }

    // no losses
    output.add(players
        .stream()
        .filter(x -> !lossCounts.containsKey(x))
        .collect(Collectors.toList()));

    // exactly one loss
    output.add(lossCounts.entrySet()
        .stream()
        .filter(a -> a.getValue() == 1)
        .map(HashMap.Entry::getKey)
        .sorted(Comparator.naturalOrder())
        .collect(Collectors.toList()));

    return output;
  }
}
