#include <iostream>

using namespace std;

void printArray(int arr[], int size)
{
  for (int i = 0; i < size; i++)
  {
    cout << arr[i] << endl;
  }
  cout << endl;
}

int main()
{
  int size = 5;

  int *Arr = new int[size]; // new keyword allocates memory in bytes

  for (int i = 0; i < size; i++)
  {
    Arr[i] = i;
  }

  printArray(Arr, size);

  int *ResizeArray = new int[size + 2];

  for (int i = 0; i < size; i++)
  {
    ResizeArray[i] = Arr[i];
  }

  delete[] Arr; // delete frees the allocated space.

  Arr = ResizeArray;

  Arr[size] = 5;
  Arr[size + 1] = 6;

  printArray(Arr, size + 2);
}