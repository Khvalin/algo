using System;
using System.Collections.Generic;
using System.IO;
using System.Text;

namespace _1234D {
  class Solution {
    const int NMAX = 'z' - 'a' + 1;
    char[] input;
    List<ulong[]> BIT = new List<ulong[]> ();

    public void ReplaceChar (int pos, char newChar) {
      var prev = this.input[pos];
      this.input[pos] = newChar;

      while (pos < this.BIT.Count) {
        this.BIT[pos][prev - 'a']--;
        this.BIT[pos][newChar - 'a']++;

        pos += (pos & -pos);
      }
    }

    private ulong[] prefixSum (int i) {
      var res = new ulong[NMAX];

      while (i != 0) {
        for (var k = 0; k < NMAX; k++) {
          res[k] += this.BIT[i][k];
        }

        i -= (i & -i);
      }

      return res;
    }

    public int CountUniqueChars (int l, int r) {
      int res = 0;

      var left = this.prefixSum (l - 1);
      var right = this.prefixSum (r);

      for (var i = 0; i < NMAX; i++) {
        if (right[i] - left[i] != 0) {
          res++;
        }
      }

      return res;
    }

    public Solution (string s) {
      s = " " + s;
      this.input = s.ToCharArray ();

      BIT = new List<ulong[]> ();
      BIT.Add (new ulong[NMAX]);

      for (int i = 1; i < s.Length; i++) {
        BIT.Add (new ulong[NMAX]);

        if (i < s.Length) {
          var ch = s[i];

          this.BIT[i][ch - 'a'] += 1;
        }
      }

      for (int i = 1; i < s.Length; i++) {
        var j = i + (i & -i);

        if (j < s.Length) {
          for (var k = 0; k < NMAX; k++) {
            this.BIT[j][k] += this.BIT[i][k];
          }
        }
      }
    }
  }

  class Program {
    static void Main (string[] args) {
      string s = Console.ReadLine ();
      Solution sol = new Solution (s);

      ulong q = ulong.Parse (Console.ReadLine ());

      for (ulong i = 0; i < q; i++) {
        var testCase = Console.ReadLine ().Split ();
        var pos = int.Parse (testCase[1]);

        if (testCase[0][0] == '1') {
          var ch = testCase[2][0];
          sol.ReplaceChar (pos, ch);
        }

        if (testCase[0][0] == '2') {
          var r = int.Parse (testCase[2]);
          var res = sol.CountUniqueChars (pos, r);

          Console.WriteLine (res);
        }
      }
    }

  }
}