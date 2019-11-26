using System;
using System.Collections.Generic;
using System.Linq;

namespace maximum_score_words_formed_by_letters {
  public class Solution {

    public class StateComparer : IEqualityComparer<IEnumerable<bool>> {
      public bool Equals (IEnumerable<bool> x, IEnumerable<bool> y) {
        if (x == null || y == null) {
          return x == y;
        }

        var res = x.Count () == y.Count ();
        var i = 0;

        while (res && i < x.Count ()) {
          res = x.ElementAt (i) == y.ElementAt (i);
          i++;
        }
        return res;
      }

      public int GetHashCode (IEnumerable<bool> obj) {
        if (obj == null) {
          return 0;
        }
        return string.Join (" ", obj.Select (b => b ? '1' : '0')).GetHashCode ();
      }
    }

    const int NMAX = 28;

    public int MaxScoreWords (string[] words, char[] letters, int[] score) {
      var wscores = new Dictionary<int, int> ();
      var wcounts = new Dictionary<int, int[]> ();
      for (int i = 0; i < words.Length; i++) {
        var w = words[i];
        int s = 0;
        int[] key = new int[NMAX];
        foreach (char ch in w) {
          var c = ch - 'a';
          s += score[c];
          key[c]++;
        }

        wscores.Add (i, s);
        wcounts.Add (i, key);
      }
      var dict = new int[NMAX];
      foreach (char ch in letters) {
        dict[ch - 'a']++;
      }

      var state = words.Select ((w) => true);
      this.memo = new Dictionary<IEnumerable<bool>, int> (new StateComparer ());

      return Solve (state, wscores, wcounts, dict);
    }

    private Dictionary<IEnumerable<bool>, int> memo = null;

    private int Solve (IEnumerable<bool> state, Dictionary<int, int> wscores, Dictionary<int, int[]> wcounts, int[] dict) {
      var r = 0;
      if (this.memo.ContainsKey (state)) {
        return memo[state];
      }

      for (int j = 0; j < wscores.Count;) {
        var ws = wscores.ElementAt (j);
        var f = true;
        for (var i = 0; i < wcounts[ws.Key].Length; i++) {
          f = f && dict[i] >= wcounts[ws.Key][i];
        }
        if (!f) {
          wscores.Remove (ws.Key);
          continue;
        }
        j++;

        var dictCopy = dict.Clone () as int[];
        for (var i = 0; i < wcounts[ws.Key].Length; i++) {
          dictCopy[i] -= wcounts[ws.Key][i];
        }

        var scoresCopy = new Dictionary<int, int> (wscores);
        scoresCopy.Remove (ws.Key);
        var next = state.Select ((b, i) => b && i != ws.Key);

        r = Math.Max (r, ws.Value + this.Solve (next, scoresCopy, wcounts, dictCopy));
      }
      this.memo[state] = r;

      return r;
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

      words = new string[] { "xxxz", "ax", "bx", "cx" };
      letters = new char[] { 'z', 'a', 'b', 'c', 'x', 'x', 'x' };
      score = new int[] { 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10 };
      r = Sol.MaxScoreWords (words, letters, score);
      Console.WriteLine (r);

      words = new string[] { "leetcode" };
      letters = new char[] { 'l', 'e', 't', 'c', 'o', 'd' };
      score = new int[] { 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0 };
      r = Sol.MaxScoreWords (words, letters, score);
      Console.WriteLine (r);
    }
  }
}