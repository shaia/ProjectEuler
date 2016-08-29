#ifndef PROJECTEULER_PR1_1_H
#define PROJECTEULER_PR1_1_H

#include <vector>
#include <numeric>

class pr1_1
{
public:
    int Solve()
    {
        std::vector<int> range(999);
        std::vector<int> filterd_range(999);

        std::iota(range.begin(), range.end(), 1);

        std::copy_if(range.begin(),range.end(),filterd_range.begin(),
                     [&](const int& number) -> bool { return (number % 3 == 0 || number %5 == 0); });
        int result = std::accumulate(filterd_range.begin(),filterd_range.end(),0);
        return result;
    }
};

#endif //PROJECTEULER_PR1_1_H
