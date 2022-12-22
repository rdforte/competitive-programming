#include <bits/stdc++.h>
#include "rapidjson/document.h" // https://rapidjson.org/md_doc_tutorial.html installed via Homebrew.

using namespace std;
using namespace rapidjson;

int compare(Value left, Value right)
{
  for (SizeType i = 0; i < max(left.Size(), right.Size()); i++)
  {
    if (i >= left.Size())
      return 1;
    if (i >= right.Size())
      return -1;

    int res = 0;

    if (left[i].IsArray() and right[i].IsNumber())
    {
      Document document;
      Document::AllocatorType &allocator = document.GetAllocator();
      Value arrVal(kArrayType);
      arrVal.PushBack(Value().SetInt(right[i].GetInt()), allocator);
      res = compare(left[i].GetArray(), arrVal.GetArray());
    }
    else if (left[i].IsNumber() and right[i].IsArray())
    {
      Document document;
      Document::AllocatorType &allocator = document.GetAllocator();
      Value arrVal(kArrayType);
      arrVal.PushBack(Value().SetInt(left[i].GetInt()), allocator);
      res = compare(arrVal.GetArray(), right[i].GetArray());
    }
    else if (left[i].IsArray() && right[i].IsArray())
    {
      res = compare(left[i].GetArray(), right[i].GetArray());
    }
    else
    {
      int leftNum = left[i].GetInt();
      int rightNum = right[i].GetInt();
      if (leftNum == rightNum)
        continue;
      return leftNum < rightNum ? 1 : -1;
    }

    if (res == 0)
      continue;
    return res;
  }

  return 0;
}

int main()
{
  freopen("input.txt", "r", stdin);
  freopen("q1output.txt", "w", stdout);

  int count = 0;
  int index = 1;

  string left;
  string right;
  while (getline(cin, left) and getline(cin, right))
  {
    const char *leftCString = left.c_str();
    const char *rightCString = right.c_str();

    Document leftDoc;
    leftDoc.Parse(leftCString);
    Document rightDoc;
    rightDoc.Parse(rightCString);

    if (compare(leftDoc.GetArray(), rightDoc.GetArray()) == 1)
    {
      count += index;
    }
    index++;
    cin.ignore();
  }

  cout << "ANSWER: " << count;
}