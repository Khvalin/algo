using System;

namespace tossstrangecoins
{
  public class Solution
  {
    public double ProbabilityOfHeads(double[] prob, int target)
    {
      var cache = new double[prob.Length + 1, target + 1];
      for (var i = 0; i < prob.Length + 1; i++)
      {
        for (var j = 0; j < target+1; j++)
        {
          cache[i, j] = -1;
        }
      }

      Func<int, int, double> solve = null;
      solve = (int ind, int t) =>
      {
        double r=1;
        if (ind >= prob.Length)
        {
          return 0;
        }

        if (cache[ind,t] > -1)
        {
          return cache[ind, t];
        }

        if (t == 0)
        {
          r = 1;
          for (var i = ind; i < prob.Length; i++)
          {
            r *= 1 - prob[i];
          }
        }
        else
        {
          r = prob[ind] * solve(ind + 1, t - 1) + (1 - prob[ind]) * solve(ind + 1, t);
        }

        cache[ind, t] = r;

        return r;
      };

      return solve(0, target);
    }
  }


  class Program
  {
    static void Main(string[] args)
    {
      var sol = new Solution();

      var prob = new double[] { 0.5, 0.5, 0.5, 0.5, 0.5 };
      var target = 0;
      Console.WriteLine(sol.ProbabilityOfHeads(prob, target));


      prob = new double[] { 1, 1, 1 };
      target = 2;
      Console.WriteLine(sol.ProbabilityOfHeads(prob, target));
    }
  }
}
