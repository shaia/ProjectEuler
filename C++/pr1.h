//
// Created by Shai Asher on 29/08/2016.
//

#ifndef PROJECTEULER_PR1_H
#define PROJECTEULER_PR1_H

class PR1
{
public:
    int Solve()
    {
        int sum=0;
        for(int i=1;i<1000;i++)
        {
            if (i % 3 == 0 || i % 5 ==0)
                sum = sum + i;
        }
        return sum;
    }
};

#endif //PROJECTEULER_PR1_H
