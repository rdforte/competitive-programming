
#include <iostream>
#include <vector>

class Solution
{
public:
    int wateringPlants(std::vector<int> &plants, int capacity)
    {
        int steps = 0;
        int cap = capacity;

        for (int i = 0; i < plants.size(); i++)
        {
            steps++;
            cap -= plants[i];

            // if on last plant break
            if (i == plants.size() - 1)
                break;

            // if can water next plant continue
            if (i < plants.size() - 1 && cap >= plants[i + 1])
            {
                continue;
            }
            else
            {
                // cant water next plant so return to water hole and refill
                steps += 2 * (i + 1);
                cap = capacity;
            }
        }
        return steps;
    }
};

int main()
{
    std::vector<int> v{2, 2, 3, 3};
    std::cout << Solution().wateringPlants(v, 5);
}