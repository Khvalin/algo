// See https://aka.ms/new-console-template for more information

int t = int.Parse(Console.ReadLine());

using var writer = new StreamWriter(Console.OpenStandardOutput());
for (var i = 0; i < t; i++)
{
    Console.ReadLine();
    var s = Console.ReadLine();
    var res = new char[s.Length - 1];

    int j = -1;
    int sum = 0;
    foreach (var ch in s)
    {
        var d = 1;
        var sign = '+';
        if (ch == '1' && sum > 0)
        {
            sign = '-';
            d = -1;
        }

        if (ch == '0')
        {
            d = 0;
        }

        sum += d;
        j++;
        if (j == 0)
        {
            continue;
        }

        res[j - 1] = sign;
    }

    writer.WriteLine(res);
}

writer.Flush();