#include <iostream>

int main()
{
    int a;
    std::cin >> a;
    // int n[a];

    // for (int i = 0; i < a; i++)
    //     std::cin >> n[i];

    for (int i = 0; i < a; i++)
    {
        int currentNumber;
        std::cin >> currentNumber;
        int sum = 0;
        while (currentNumber / 10 >= 1)
        {
            int furthestRightNumber = currentNumber % 10;
            sum += furthestRightNumber;
            currentNumber -= furthestRightNumber;
            currentNumber /= 10;
        }
        std::cout << (sum + currentNumber) << std::endl;
    }
}