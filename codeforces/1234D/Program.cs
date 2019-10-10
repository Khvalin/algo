using System;
using System.Collections.Generic;
using System.IO;
using System.Text;

namespace _1234D
{
  class Solution
  {
    const int NMAX = 'z' - 'a' + 1;
    char[] input;
    SortedSet<ulong>[] indices = new SortedSet<ulong>[NMAX];

    public void ReplaceChar(ulong pos, char newChar)
    {
      char oldChar = this.input[pos];
      this.input[pos] = newChar;

      if (oldChar == newChar)
      {
        return;
      }

      this.indices[oldChar - 'a'].Remove(pos);
      this.indices[newChar - 'a'].Add(pos);
    }

    public int CountUniqueChars(ulong l, ulong r)
    {
      int res = 0;
      foreach (var s in this.indices)
      {
        if (s != null && s.GetViewBetween(l, r).Count > 0)
        {
          res++;
        }
      }

      return res;
    }

    public Solution(string s)
    {
      this.input = s.ToCharArray();

      for (int i = 0; i < s.Length; i++)
      {
        char ch = s[i];

        int j = ch - 'a';
        if (indices[j] == null)
        {
          indices[j] = new SortedSet<ulong>();
        }

        indices[j].Add((ulong)i);
      }

      Console.Write(String.Join(' ', indices[1]));
    }
  }

  class Program
  {
    static void Main(string[] args)
    {
      string s = Console.ReadLine();
      Solution sol = new Solution(s);

      ulong q = ulong.Parse(Console.ReadLine());

      for (ulong i = 0; i < q; i++)
      {
        var testCase = Console.ReadLine().Split();
        var pos = ulong.Parse(testCase[1]);

        if (testCase[0][0] == '1')
        {
          var ch = testCase[2][0];
          sol.ReplaceChar(pos - 1, ch);
        }

        if (testCase[0][0] == '2')
        {
          var r = ulong.Parse(testCase[2]);
          var res = sol.CountUniqueChars(pos - 1, r - 1);
          Console.WriteLine(res);
        }
      }
    }

  }
}

