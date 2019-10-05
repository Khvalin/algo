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
        arr[i] = UInt64.Parse (s[i]);
      }

      return arr.AsEnumerable ();
    }

    static int Solve (IEnumerable<ulong> input) {
      HashSet<ulong> res = new HashSet<ulong> ();

      if (input.Count () == 0) {
        return 0;
      }

      ulong min = (input.ElementAt (0));
      foreach (ulong n in input) {
        if (n < min) {
          min = n;
        }
      }

      res.Add (min);
      Console.WriteLine (min);

      for (ulong i = 2; i <= min / 2; i++) {
        if (min % i == 0) {
          res.Add (i);
        }
      }

      foreach (ulong n in input) {
        if (n == min) {
          continue;
        }

        //        Console.WriteLine (String.Join (' ', res));

        foreach (ulong m in res) {
          if (n % m > 0) {
            res.Remove (m);
          }
        }
      }

      res.Add (1);

      return res.Count ();
    }

    static void Main (string[] args) {
      IEnumerable<ulong> input = Read ();

      Console.WriteLine (Solve (input));
    }
  }
}