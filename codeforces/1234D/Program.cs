using System;
using System.IO;
using System.Text;

namespace _1234D
{
  class Solution{
    string input;

    public Solution(string s) {
      this.input = s;
    }
  }

  class Program
  {
    static void Main(string[] args)
    {
      using (Stream stdin = Console.OpenStandardInput())
      {
        string s = ReadStringFromStream(stdin);
        ulong q = ReadUlongFromStream(stdin);

        for (ulong i = 0; i < q;i++){
          ulong c  = ReadUlongFromStream(stdin);
          ulong pos  = ReadUlongFromStream(stdin);

          if (c == 1) {
            char ch = ReadCharFromStream(stdin);
          } else {
            ulong r = ReadUlongFromStream(stdin);
          }
        }
      }
    }

        static string ReadStringFromStream(Stream stream)
    {
      StringBuilder sb = new StringBuilder();

      while (stream.CanRead)
      {
        int b = stream.ReadByte();
        if (b == -1)
        {
          break;
        }

        char ch = (char)b;

        if (ch == '\n')
        {
          break;
        }

        sb.Append(ch);
      }

      return sb.ToString();
    }

    static char ReadCharFromStream(Stream stream)
    {
      char ch = (char) stream.ReadByte() ;

      return ch;
    }

    static ulong ReadUlongFromStream(Stream stream)
    {
      ulong n = 0;
      bool f = false;

      while (stream.CanRead)
      {
        int b = stream.ReadByte();
        int d = b - '0';

        if (d < 0 || d > 9)
        {
          if (f)
          {
            break;
          }

          continue;
        }
        f = true;

        n *= 10;
        n += (ulong)d;
      }

      return n;
    }
  }
}

