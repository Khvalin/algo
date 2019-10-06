using System;
using System.Collections.Generic;

namespace count_vowels_permutation {

  public class Solution {
    const int MOD = 1000000007;

    public int CountVowelPermutation (int n) {
      var count = new Dictionary<char, ulong> { { 'a', 1 },
          { 'e', 1 },
          { 'i', 1 },
          { 'o', 1 },
          { 'u', 1 },
        };

      var next = new Dictionary<char, ulong> (count);
      for (int i = 1; i < n; i++) {
        next['a'] = (count['e'] + count['i'] + count['u']) % MOD;
        next['e'] = (count['a'] + count['i']) % MOD;
        next['i'] = (count['e'] + count['o']) % MOD;
        next['o'] = (count['i']);
        next['u'] = (count['i'] + count['o']) % MOD;

        count['a'] = next['a'];
        count['e'] = next['e'];
        count['i'] = next['i'];
        count['o'] = next['o'];
        count['u'] = next['u'];
      }

      return (int) ((count['a'] + count['e'] + count['i'] + count['o'] + count['u']) % MOD);
    }
  }

  class Program {
    static void Main (string[] args) {
      Solution s = new Solution ();
      Console.WriteLine (s.CountVowelPermutation (2));
      Console.WriteLine (s.CountVowelPermutation (5));
      Console.WriteLine (s.CountVowelPermutation (20000));
    }
  }
}