using System;
using System.Linq;

public static class PR11
{
    public static int Solve()
    {
        int ans = Enumerable.Range(1,1000).Where( i=> i%3==0 || i%5==0 ).Sum();
        return ans;
    }
}