package org.pateason.leetcode;

import java.util.*;
import java.util.stream.Collectors;

/**
 * @see <a href="https://leetcode.com/problems/top-k-frequent-words/">Leetcode</a>
 *
 * Given an array of strings words and an integer k,
 * return the k most frequent strings.
 *
 * Return the answer sorted by the frequency from highest
 * to lowest. Sort the words with the same frequency by
 * their lexicographical order.
 */
public class TopKFrequentWords {
  public static List<String> topKFrequent(String[] words, int k) {
    List<String> output = new ArrayList<>();

    // early return if there is nothing in the stack
    if (words.length == 0) {
      return output;
    }

    // get counts for each word, in order of appearance
    LinkedHashMap<String, Integer> wordCounts = new LinkedHashMap<>();
    for (String word : words) {
      wordCounts.put(
          word,
          wordCounts.getOrDefault(word, 0) + 1
      );
    }

    // group by count, sort the outputs
    TreeMap<Integer, List<String>> wordCountsMap = new TreeMap<>(Collections.reverseOrder());
    wordCountsMap.putAll(
      wordCounts
        .keySet()
        .stream()
        .sorted()
        .collect(Collectors.groupingBy(wordCounts::get))
    );

    // generate output array
    output = wordCountsMap
      .values()
      .stream()
      .flatMap(List::stream)
      .collect(Collectors.toList());

    int outputSize = Math.min(output.size(), k);
    return output.subList(0, outputSize);
  }
}
