using System;
using System.Linq;
using System.Collections.Generic;

namespace _1256D
{
  class Program
  {

    static IEnumerable<bool> Solve(List<bool> input, long k)
    {
      var L = input.Count();

      var left = L;
      var zeroes = new List<int>();

      for (var i = L - 1; i >= 0; i--)
      {
        if (input.ElementAt(i))
        {
          left = i;
        }
        else
        {
          zeroes.Add(i);
        }
      }

      var z = zeroes.Count - 1;

      while (k > 0 && z >= 0 && left < L)
      {
        var i = zeroes[z];
        z--;

        if (i <= left)
        {
          continue;
        }

        if (k >= i - left)
        {
          input[i] = true;
          input[left] = false;

          k -= i - left;

          while (left < L && !input[left])
          {
            left++;
          }
        }
        else
        {
          var r = i;
          var j = i;

          while (k > 0 && j >= left)
          {
            j--;
            if (input[j])
            {
              r = j;
            }

            k--;
          }

          input[i] = true;
          input[r] = false;
        }
      }

      return input;
    }

    static void Main(string[] args)
    {
      var q = int.Parse(Console.ReadLine());

      for (var i = 0; i < q; i++)
      {
        var line = Console.ReadLine().Split(new char[] { ' ' });
        var n = int.Parse(line[0]);
        var k = long.Parse(line[1]);

        var s = Console.ReadLine();
        var input = s.Select(ch => ch.Equals('1'));

        input = Solve(new List<bool>(input), k);

        Console.WriteLine(String.Join("", input.Select(b => b ? 1 : 0)));
      }
    }
  }
}
