using System;
using System.Collections.Generic;

namespace maximum_score_words_formed_by_letters {
  public class Solution {
    public int MaxScoreWords (string[] words, char[] letters, int[] score) {
      var scores = new Dictionary<int, int> ();
      foreach (var w in words) {
        // scores.Add ()
      }

      return 0;
    }
  }

  class Program {
    static void Main (string[] args) {
      var Sol = new Solution ();
      var words = new string[] { "dog", "cat", "dad", "good" };
      var letters = new char[] { 'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o' };
      var score = new int[] { 1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 };

      var r = Sol.MaxScoreWords (words, letters, score);
      Console.WriteLine (r);
    }
  }
}