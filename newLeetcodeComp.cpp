#include <iostream>
#include <filesystem>
#include <fstream>
#include <iostream>
#include <string>

using namespace std;
namespace fs = std::filesystem;

int main()
{
  int competitionNumber;

  cout << "🏆\033[92m MAKE LEET_CODE COMP \033[0m🏆\n";
  cout << "what is the competition number?\n";
  cin >> competitionNumber;

  if (competitionNumber)
  {
    string directory = to_string(competitionNumber);
    fs::create_directories("./LeetCode/Competitions/" + directory);

    ofstream file;

    for (int i = 1; i <= 4; i++)
    {

      file.open("./LeetCode/Competitions/" + directory + "/" + "q" + to_string(i) + ".cpp");
      file << "#include <iostream>\n";
      file << "#include <queue>\n";
      file << "#include <stack>\n";
      file << "#include <vector>\n";
      file << "#include <set>\n";
      file << "#include <algorithm>\n";
      file << "#include <utility>\n";
      file << "#include <unordered_set>\n";
      file << "#include <unordered_map>\n";
      file << "\n";
      file << "using namespace std;\n";
      file << "\n";
      file.close();
    }

    cout << "created comp " + directory;
  }
  else
  {
    cout << "\nLooks like this is not a valid competition number.";
  }
}