using System;
using System.Collections.Generic;
using System.Text.Json;

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
    private int[][] ships;
    public int count;

    public Sea (int[][] ships) {
      this.ships = ships;
    }

    public bool HasShips (int[] topRight, int[] bottomLeft) {
      this.count++;
      foreach (var ship in this.ships) {
        if (ship[0] >= bottomLeft[0] && ship[0] <= topRight[0] && ship[1] >= bottomLeft[1] && ship[1] <= topRight[1]) {
          return true;
        }
      }

      return false;
    }
  }

  class Solution {
    public int CountShips (Sea sea, int[] topRight, int[] bottomLeft) {
      var r = 0;

      List<int[]> points = new List<int[]> (this.Split (topRight, bottomLeft));

      var cache = new System.Collections.Generic.Dictionary<string, bool> ();

      while (points.Count > 0 && r < 10) {
        var ind = 0;
        var p = points[ind];
        points.RemoveAt (ind);
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

          var split = this.Split (topRight, bottomLeft);

          points.AddRange (split);
          //   Console.WriteLine ("{0} {1} {2}", points.Count, JsonSerializer.Serialize (p), JsonSerializer.Serialize (split));
        }
      }

      return r;
    }

    private IEnumerable<int[]> Split (int[] topRight, int[] bottomLeft) {
      var dx = Math.Abs (topRight[0] - bottomLeft[0]);
      var dy = Math.Abs (topRight[1] - bottomLeft[1]);

      if (dx >= dy && dx <= 1) {
        if (dx == 0 || dy == 0) {
          return (new int[][] {
            //new int[] { bottomLeft[0], bottomLeft[1], bottomLeft[0], bottomLeft[1] },
            //  new int[] { topRight[0], topRight[1], topRight[0], topRight[1] },
            new int[] { bottomLeft[0], topRight[1], bottomLeft[0], topRight[1] },
              new int[] { topRight[0], bottomLeft[1], topRight[0], bottomLeft[1], }
          });
        }

        return (new int[][] {
          new int[] { bottomLeft[0], bottomLeft[1], bottomLeft[0], bottomLeft[1] },
            new int[] { topRight[0], topRight[1], topRight[0], topRight[1] },
            new int[] { bottomLeft[0], topRight[1], bottomLeft[0], topRight[1] },
            new int[] { topRight[0], bottomLeft[1], topRight[0], bottomLeft[1], }
        });
      }

      int[] mid = new int[] { bottomLeft[0], bottomLeft[1] };

      mid[0] += dx >> 1;
      mid[1] += dy >> 1;

      if (dx > 250 && mid[0] % 2 == 0) {
        mid[0] += 3;
      }

      if (dy > 250 && mid[1] % 2 == 0) {
        mid[1] += 3;
      }

      return (new int[][] {
        new int[] { bottomLeft[0], bottomLeft[1], mid[0] - mid[0] % 2, mid[1] - mid[1] % 2 },
          new int[] { bottomLeft[0], mid[1], mid[0], topRight[1] },
          new int[] { mid[0] + mid[0] % 2, mid[1] + mid[1] % 2, topRight[0], topRight[1] },
          new int[] { mid[0], bottomLeft[1], topRight[0], mid[1] }
      });
    }
  }

  class Program {
    static void Test51 () {
      var sol = new Solution ();
      var ships = JsonSerializer.Deserialize<int[][]> ("[[53,118],[373,246],[121,625],[805,303],[383,532],[982,664],[453,122],[347,106]]");
      var sea = new Sea (ships);

      var topRight = new int[] { 1000, 1000 };
      var bottomLeft = new int[] { 0, 0 };
      Console.WriteLine (sol.CountShips (sea, topRight, bottomLeft) + " " + sea.count);
    }

    static void Test53 () {
      var sol = new Solution ();
      var ships = JsonSerializer.Deserialize<int[][]> ("[[556,290],[286,960],[728,419],[770,159],[743,306],[674,140],[37,232],[515,544],[883,449]]");
      var sea = new Sea (ships);

      var topRight = new int[] { 1000, 1000 };
      var bottomLeft = new int[] { 0, 0 };
      Console.WriteLine (sol.CountShips (sea, topRight, bottomLeft) + " " + sea.count);
    }

    static void Test56 () {
      var sol = new Solution ();
      var ships = JsonSerializer.Deserialize<int[][]> ("[[334,428],[137,522],[772,606],[143,341],[449,783],[784,453],[958,782],[682,500],[845,228],[125,238]]");
      var sea = new Sea (ships);

      var topRight = new int[] { 1000, 1000 };
      var bottomLeft = new int[] { 0, 0 };
      Console.WriteLine (sol.CountShips (sea, topRight, bottomLeft) + " " + sea.count);
    }

    static void Test56_1 () {
      var sol = new Solution ();
      var ships = JsonSerializer.Deserialize<int[][]> ("[[658,635],[617,840],[514,853],[507,332]]");
      var sea = new Sea (ships);

      var topRight = new int[] { 1000, 1000 };
      var bottomLeft = new int[] { 0, 0 };
      Console.WriteLine ("{0} {1} {2}", 4, sol.CountShips (sea, topRight, bottomLeft), sea.count);
    }

    static void Main (string[] args) {
      // Test51 ();
      // Test53 ();
      Test56_1 ();
      Test56 ();
    }
  }
}