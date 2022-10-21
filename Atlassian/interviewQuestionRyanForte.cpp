#include <bits/stdc++.h>

using namespace std;

// Space Complexity O(n)
class FileSizes
{
  set<pair<int, string>, greater<pair<int, string>>> orderedFileBySizes;
  int totalSize = 0;

  // Time Complexity O(n)
  // Space Complexity O(n)
public:
  FileSizes(vector<pair<string, int>> filesSizes)
  {
    unordered_map<string, int> files;

    for (auto file : filesSizes)
    {
      if (file.first == "")
      {
        totalSize += file.second;
        continue;
      }

      if (files[file.first])
      {
        files[file.first] += file.second;
      }
      else
      {
        files[file.first] += file.second;
      }
      totalSize += file.second;
    }

    for (auto file : files)
    {
      orderedFileBySizes.insert({file.second, file.first});
    }
  }

  // Time Complexity O(1);
  int findTotalSizeOfAllFiles()
  {
    return totalSize;
  }

  // Time Complexity O(n)
  // Space Complexity O(n)
  vector<string> findTopFileByCollections(int n)
  {
    vector<string> collections;

    for (auto file : orderedFileBySizes)
    {
      if (n <= 0)
        break;
      collections.push_back(file.second);
      n--;
    }

    return collections;
  };
};

int main()
{
  vector<pair<string, int>> files = {
      {"col1", 10},
      {"col1", 10},
      {"col2", 10},
      {"col1", 10},
      {"col2", 70},
      {"", 500}, // file not belonging to a collection
      {"col3", 1000}};

  FileSizes *f = new FileSizes(files);

  // Check to see if all files add to the expected value.

  int totalSizeOfAllFiles = f->findTotalSizeOfAllFiles();
  int expectedTotalFileSize = 1610;
  if (totalSizeOfAllFiles != expectedTotalFileSize)
  {
    throw logic_error("expected total file size "s + to_string(1610) + " except got "s + to_string(totalSizeOfAllFiles));
  }
  cout << "total size of all files should be 1610\n";

  // Check top collections

  vector<string> topCollections = f->findTopFileByCollections(2);
  int expectedCollectionSize = 2;
  if (topCollections.size() != expectedCollectionSize)
  {
    throw logic_error("expected collection size "s + to_string(expectedCollectionSize) + " except got "s + to_string(topCollections.size()));
  }
  cout << "collection size should be size 2\n";

  vector<string> expectedCollections = {
      "col3",
      "col2"};

  for (int i = 0; i < topCollections.size(); i++)
  {
    if (expectedCollections[i] != topCollections[i])
    {
      throw logic_error("expected collection "s + expectedCollections[i] + " except got "s + topCollections[i]);
    }
  }
  cout << "expect collections to be in order of col3, col2\n";
}