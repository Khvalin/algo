using System;
using System.Collections.Generic;

//number-of-ships-in-a-rectangle
namespace number_of_ships_in_a_rectangle {

  /**
   * // This is Sea's API interface.
   * // You should not implement it, or speculate about its implementation
   * class Sea {
   *     public bool HasShips(int[] topRight, int[] bottomLeft);
   * }
   */

  class Sea {
    public bool HasShips (int[] topRight, int[] bottomLeft) {
      return true;
    }
  }

  class Solution {
    public int CountShips (Sea sea, int[] topRight, int[] bottomLeft) {
      var r = 0;

      List<int[]> points = new List<int[]> (this.Split (topRight, bottomLeft));

      var cache = new System.Collections.Generic.Dictionary<string, bool> ();

      while (points.Count > 0) {
        var p = points[0];
        points.RemoveAt (0);
        var key = String.Join (' ', p);
        if (cache.ContainsKey (key)) {
          continue;
        }

        cache[key] = true;

        topRight[0] = p[2];
        topRight[1] = p[3];
        bottomLeft[0] = p[0];
        bottomLeft[1] = p[1];

        if (sea.HasShips (topRight, bottomLeft)) {
          if (p[0] == p[2] && p[1] == p[3]) {
            r++;
            continue;
          }

          points.AddRange (this.Split (topRight, bottomLeft));
        }
      }

      return r;
    }

    private IEnumerable<int[]> Split (int[] topRight, int[] bottomLeft) {
      int[] mid = new int[] { bottomLeft[0], bottomLeft[1] };

      mid[0] += (Math.Abs (topRight[0] - bottomLeft[0]) >> 1);
      mid[1] += (Math.Abs (topRight[1] - bottomLeft[1]) >> 1);

      return (new int[][] {
        new int[] { bottomLeft[0], bottomLeft[1], mid[0], mid[1] },
          new int[] { bottomLeft[0], mid[1], mid[0], topRight[1] },
          new int[] { mid[0], mid[1], topRight[0], topRight[1] },
          new int[] { mid[0], bottomLeft[1], topRight[0], mid[1] }
      });
    }
  }

  class Program {
    static void Main (string[] args) {
      var sol = new Solution ();
      var sea = new Sea ();

      var topRight = new int[] { 4, 4 };
      var bottomLeft = new int[] { 0, 0 };

      Console.WriteLine (sol.CountShips (sea, topRight, bottomLeft));
    }
  }
}