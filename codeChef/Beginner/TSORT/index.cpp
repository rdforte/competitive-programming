#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

// using insertion sort
// int main()
// {
//     int input;
//     vector<int> vec;

//     while (cin >> input)
//     {
//         vec.push_back(input);
//         for (int i = 1; i < vec.size(); i++)
//         {
//             int crnt = vec[i];
//             int prevIndex = i - 1;
//             while (prevIndex >= 0 && vec[prevIndex] > crnt)
//             {
//                 vec[prevIndex + 1] = vec[prevIndex];
//                 prevIndex--;
//             };
//             vec[prevIndex + 1] = crnt;
//         }
//     }

//     for (int i = 0; i < vec.size(); i++)
//         cout << vec[i] << endl;
// }

/**
 * When we use the sort function to sort an array our arguments will look a bit different then when we use it on a vector for example. 
 * In the example above when we pass in intArray as an argument we are telling the function to start the sort at the beginning of the array. 
 * If we wanted it to start the sort at the second element of the array we would do sort(intArray + 1, intArray + SIZE);. 
 * So when we do intArray + SIZE for the second argument we are telling the array to sort up to the last element in the array.
*/
int main()
{
    int t;
    cin >> t;
    int arr[t];
    for (int i = 0; i < t; i++)
    {
        cin >> arr[i];
    }
    sort(arr, arr + t); // sort(intArray + 1, intArray + SIZE);
    for (int i = 0; i < t; i++)
    {
        cout << arr[i] << endl;
    }
}