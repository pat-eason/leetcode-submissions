package org.pateason.leetcode;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.List;

class TopKFrequentWordsTest {

  @Test
  void topKFrequent_2() {
    int topFrequentWords = 2;
    String[] inputWords = new String[] { "i", "love", "leetcode", "i", "love", "coding" };

    List<String> expectedOutput = new ArrayList<>() {{
      add("i");
      add("love");
    }};

    List<String> result = TopKFrequentWords.topKFrequent(inputWords, topFrequentWords);
    Assertions.assertEquals(expectedOutput, result);
  }

  @Test
  void topKFrequent_4() {
    int topFrequentWords = 4;
    String[] inputWords = new String[] { "the","day","is","sunny","the","the","the","sunny","is","is" };

    List<String> expectedOutput = new ArrayList<>() {{
      add("the");
      add("is");
      add("sunny");
      add("day");
    }};

    List<String> result = TopKFrequentWords.topKFrequent(inputWords, topFrequentWords);
    Assertions.assertEquals(expectedOutput, result);
  }

//  @Test
//  void topKFrequent_0() {
//  }
}