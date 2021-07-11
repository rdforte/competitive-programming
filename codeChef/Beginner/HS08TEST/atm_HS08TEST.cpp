#include <iostream>
#include <iomanip> // this is required in order to use setPrecision

int main()
{
    int w;
    float b;

    std::cin >> w >> b;

    // setprecision makes it so we can specify the number of decimal places
    // fixed means floating-point values will be written in fixed point notations.

    b = w % 5 != 0 || w > b - 0.5 ? b : b - w - 0.5;
    std::cout << std::fixed << std::setprecision(2) << b;
    return 0;
};

// #include <iostream>
// #include <iomanip>

// int main()
// {
//     int w;
//     float b;

//     std::cin >> w >> b;

//     b = w % 5 == 0 && w <= b - 0.5 ? b - w - 0.5 : b;
//     std::cout << std::fixed << std::setprecision(2) << b;
//     return 0;
// };