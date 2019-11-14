using System;
using System.Linq;
using System.Collections.Generic;

namespace _1256D
{
  class Program
  {

    static bool[] Solve(bool[] input, long k)
    {
      var L = input.Length;

      var left = L;
      var zeroes = new List<int>();

      for (var i = 0; i < L; i++)
      {
        if (input[i])
        {
          if (left >= L)
          {
            left = i;
          }
        }
        else
        {
          zeroes.Add(i);
        }


      }

      if (left >= L)
      {
        return input;
      }



      return input;
    }

    static void Main(string[] args)
    {
      var q = int.Parse(Console.ReadLine());

      for (var i = 0; i < q; i++)
      {
        var line = Console.ReadLine().Split(" ");
        var n = int.Parse(line[0]);
        var k = long.Parse(line[1]);

        var input = new bool[n];
        var s = Console.ReadLine();

        for (var j = 0; j < n; j++)
        {
          input[j] = s[j] == '1';
        }

        input = Solve(input, k);

        Console.WriteLine(String.Join("", input.Select(b => b ? 1 : 0)));
      }
    }
  }
}
