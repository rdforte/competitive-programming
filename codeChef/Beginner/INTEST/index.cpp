#include <iostream>

int main()
{
    int i;
    int d;

    std::cin >> i >> d;

    int n = 0;

    while (i--)
    {
        long long x;
        std::cin >> x;
        if (x % d == 0)
            n++;
    }

    std::cout << n;
}