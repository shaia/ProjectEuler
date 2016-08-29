using System;
public static class PR12
{
    public static int Solve()
    {
        Func<int, int> sum = n => { var k = (999) / n; return n * k * (k + 1) / 2; };
        return sum(3) + sum(5) - sum(15);
    }
}