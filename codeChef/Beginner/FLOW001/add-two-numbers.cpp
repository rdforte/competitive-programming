#include <iostream>
#include <vector>

// My Solution

// int main()
// {
//     int numberOfAdditions;

//     std::cin >> numberOfAdditions;

//     std::vector<int> additions;

//     for (int currentAdditions = 1; currentAdditions <= numberOfAdditions; currentAdditions++)
//     {
//         int a;
//         int b;
//         std::cin >> a >> b;
//         additions.push_back(a + b);
//     }

//     for (int i = 0; i < additions.size(); i++)
//     {
//         std::cout << additions[i] << std::endl;
//     }
// }

// Other persons
int main()
{
    int t;
    std::cin >> t;
    while (t--) // this is going to be false when t reaches 0. So we subtract 1 from t and then run the loop
    {

        int a, b;
        std::cin >> a >> b;
        std::cout << a + b << std::endl;
    }
}