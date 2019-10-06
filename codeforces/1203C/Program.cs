using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;

namespace _1203C {
  class Program {
    static IEnumerable<ulong> Read () {
      ulong n = UInt64.Parse (Console.ReadLine ());

      ulong[] arr = new ulong[n];
      string[] s = Console.ReadLine ().Split (' ');

      for (ulong i = 0; i < n; i++) {
        ulong a = 0;
        foreach (char ch in s[i]) {
          a *= 10;
          a += (ulong) (ch - '0');
        }
        arr[i] = a;
      }

      return arr.AsEnumerable ();
    }

    static ulong GCD (ulong a, ulong b) {
      if (a == 0)
        return b;
      return GCD (b % a, a);
    }

    static int Solve (IEnumerable<ulong> input) {
      HashSet<ulong> res = new HashSet<ulong> ();

      if (input.Count () == 0) {
        return 0;
      }

      ulong min = (input.ElementAt (0));
      ulong max = (input.ElementAt (0));
      foreach (ulong n in input) {
        min = Math.Min (min, n);
        max = Math.Max (max, n);
      }

      var g = GCD (min, max);

      foreach (ulong n in input) {
        g = GCD (n, g);
      }

      for (ulong i = 2; i <= Math.Sqrt (g); i++) {
        if (g % i == 0) {
          res.Add (i);
          res.Add (g / i);
        }
      }

      res.Add (1);
      res.Add (g);

      return res.Count ();
    }

    static void Main (string[] args) {
      IEnumerable<ulong> input = Read ();

      Console.WriteLine (Solve (input));
    }
  }
}