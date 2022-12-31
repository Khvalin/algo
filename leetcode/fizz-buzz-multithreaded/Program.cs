using System.Threading;

void main()
{
    FizzBuzz instance = new FizzBuzz(10);

    var actions = new Action[]
    {
        (() => instance.Number(Console.WriteLine)),
        (() => instance.Buzz(() => Console.WriteLine("buzz"))),
        (() => instance.Fizz(() => Console.WriteLine("fizz"))),
        (() => instance.Fizzbuzz(() => Console.WriteLine("fizzbuzz"))),
    };


    Parallel.Invoke(actions);

    Console.WriteLine("Finish");
}


main();

public class FizzBuzz
{
    private int n;
    private int c = 0;
    private int ind = 0;
    private List<AutoResetEvent> _pauseEvents = new List<AutoResetEvent>();
    private Object l = new Object();


    public FizzBuzz(int n)
    {
        this.n = n;

        this._pauseEvents.Add(new AutoResetEvent(true));
        this._pauseEvents.Add(new AutoResetEvent(false));
        this._pauseEvents.Add(new AutoResetEvent(false));
        this._pauseEvents.Add(new AutoResetEvent(false));
    }

    private void Inc()
    {
        Interlocked.Increment(ref this.c);
    }

    private void Recalc()
    {
        lock (this.l)
        {
            this._pauseEvents[0].Reset();
            this._pauseEvents[1].Reset();
            this._pauseEvents[2].Reset();
            this._pauseEvents[3].Reset();

            this.ind++;
            this._pauseEvents[(this.ind) % 4].Set();
        }
    }

    // printFizz() outputs "fizz".
    public void Fizz(Action printFizz)
    {
        while (this.c <= this.n)
        {
            if (this.c > this.n)
            {
                break;
            }

            if (this.c % 3 == 0 && this.c % 5 != 0)
            {
                printFizz.Invoke();
            }

            this.Recalc();
            this._pauseEvents[1].WaitOne();
        }

        this.Recalc();
    }

    // printBuzzz() outputs "buzz".
    public void Buzz(Action printBuzz)
    {
        while (this.c <= this.n)
        {
            if (this.c > this.n)
            {
                break;
            }

            if (this.c % 3 != 0 && this.c % 5 == 0)
            {
                printBuzz.Invoke();
            }

            this.Recalc();
            this._pauseEvents[2].WaitOne();
        }

        this.Recalc();
    }

    // printFizzBuzz() outputs "fizzbuzz".
    public void Fizzbuzz(Action printFizzBuzz)
    {
        while (this.c <= this.n)
        {
            if (this.c > this.n)
            {
                break;
            }

            if (this.c > 0 && this.c % 3 == 0 && this.c % 5 == 0)
            {
                printFizzBuzz.Invoke();
            }

            this.Recalc();
            this._pauseEvents[3].WaitOne();
        }

        this.Recalc();
    }

    // printNumber(x) outputs "x", where x is an integer.
    public void Number(Action<int> printNumber)
    {
        while (this.c <= this.n)
        {
            this.Inc();

            if (this.c > this.n)
            {
                break;
            }

            if (this.c % 3 > 0 && this.c % 5 > 0)
            {
                printNumber.Invoke(this.c);
                continue;
            }

            this.Recalc();
            this._pauseEvents[0].WaitOne();
        }

        this.Recalc();
    }
}
