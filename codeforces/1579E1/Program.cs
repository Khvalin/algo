using System;
using System.IO;
using System.Collections.Generic;
using System.Linq;

namespace _1579E1
{
    class Program
    {
        static IEnumerable<int> Solve(IEnumerable<int> arr) {
            List<int> left =new List<int>(); 
            List<int> right =new List<int>(); 

            
            for (int i = 0; i < arr.Count(); i++) {
                var n = arr.ElementAt(i);
                if(i > 0 && right[right.Count - 1] < n) {
                    left.Add(n);
                    continue;
                }

                right.Add(n);
            }

            left.Reverse();
            List<int> res = left;
            res.AddRange(right);
            res.Reverse();

            return res;
        }

        static void Main(string[] args)
        {
            string input = Console.In.ReadToEnd();
            string[] lines = input.Split(Environment.NewLine, StringSplitOptions.RemoveEmptyEntries);
             
            for (int i = 2; i < lines.Length; i+=2) {
                IEnumerable<int> arr = lines[i].Split(' ').Select(s => Convert.ToInt32(s)) ;
                arr = Solve(arr);

                Console.WriteLine(String.Join(' ', arr));
            }
        }
    }
}
