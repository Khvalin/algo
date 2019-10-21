using System;

namespace maximum_profit_in_job_scheduling {
  public class Solution {

    public struct job {
      public job (int s, int e, int p) {
        this.start = s;
        this.end = e;
        this.profit = p;
      }

      public int start;
      public int end;
      public int profit;
    }

    public int JobScheduling (int[] startTime, int[] endTime, int[] profit) {
      var L = startTime.Length;

      var jobs = new job[L];

      var memo = new int[startTime.Length];
      for (var i = 0; i < memo.Length; i++) {
        memo[i] = -1;
        jobs[i] = new job (startTime[i], endTime[i], profit[i]);
      }

      Array.Sort (jobs, new Comparison<job> ((job j1, job j2) => j1.start.CompareTo (j2.start)));

      Func<int, int> calc = null;
      calc = (int ind) => {
        if (ind >= L) {
          return 0;
        }

        if (memo[ind] > -1) {
          return memo[ind];
        }

        var r = jobs[ind].profit;
        var i = ind + 1;

        while (i < L && jobs[i].start < jobs[ind].end) { i++; }

        if (i < L) {
          r += calc (i);
        }

        r = Math.Max (r, calc (ind + 1));

        memo[ind] = r;

        return r;
      };

      return calc (0);
    }
  }

  class Program {
    static void Main (string[] args) {
      var s = new Solution ();

      var input1 = new int[] { 1, 2, 3, 3 };
      var input2 = new int[] { 3, 4, 5, 6 };
      var input3 = new int[] { 50, 10, 40, 70 };
      Console.WriteLine (s.JobScheduling (input1, input2, input3));

      input1 = new int[] { 1, 2, 3, 4, 6 };
      input2 = new int[] { 3, 5, 10, 6, 9 };
      input3 = new int[] { 20, 20, 100, 70, 60 };
      Console.WriteLine (s.JobScheduling (input1, input2, input3));

      input1 = new int[] { 1, 1, 1 };
      input2 = new int[] { 2, 3, 4 };
      input3 = new int[] { 5, 6, 4 };
      Console.WriteLine (s.JobScheduling (input1, input2, input3));

      input1 = new int[] { 6, 15, 7, 11, 1, 3, 16, 2 };
      input2 = new int[] { 19, 18, 19, 16, 10, 8, 19, 8 };
      input3 = new int[] { 2, 9, 1, 19, 5, 7, 3, 19 };
      Console.WriteLine (s.JobScheduling (input1, input2, input3));

    }
  }
}