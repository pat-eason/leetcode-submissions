package org.pateason.leetcode;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;

public class GroupAnagrams {
  public static List<List<String>> groupAnagrams(String[] strs) {
    HashMap<String, List<String>> map = new HashMap<>();
    for (String str : strs) {
      char[] chars = str.toCharArray();
      Arrays.sort(chars);
      String sortKey = new String(chars);

      List<String> anList = map.get(sortKey);
      if (anList == null) {
        anList = new ArrayList<>();
      }
      anList.add(str);
      map.put(sortKey, anList);
    }

    return new ArrayList<>(map.values());
  }
}
