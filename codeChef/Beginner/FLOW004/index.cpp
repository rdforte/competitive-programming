#include <iostream>

int findFirstDigit(int n)
{
    while (n >= 10)
        n /= 10;

    return n;
}

int main()
{

    int n;
    std::cin >> n;

    for (int i = 0; i < n; i++)
    {
        int num;
        std::cin >> num;
        std::cout << findFirstDigit(num) + num % 10 << "\n";
    }

    return 0;
}