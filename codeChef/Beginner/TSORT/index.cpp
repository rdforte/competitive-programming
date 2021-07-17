#include <iostream>
#include <vector>
using namespace std;

int main()
{
    int input;
    vector<int> vec;

    while (cin >> input)
        vec.push_back(input);

    for (int i = 1; i < vec.size(); i++)
    {
        int crnt = vec[i];
        int prevIndex = i - 1;
        while (prevIndex >= 0 && vec[prevIndex] > crnt)
        {
            vec[prevIndex + 1] = vec[prevIndex];
            prevIndex--;
        };
        vec[prevIndex + 1] = crnt;
    }

    for (int i = 0; i < vec.size(); i++)
        cout << vec[i] << endl;
}