#include <bits/stdc++.h>

using namespace std;

// Sliding Window
// time complexity = the window always moves forward with each iteration. O(n) where n is the length of the string.
// space complexity = we have an array which keeps track of all uppercase characters. We can call this m. the space
// is then O(m).

// Things to note here. If possible use an array over a map because it is ~100x faster. A map can have some large constant
// values under the hood which can increase the overall time of our algorithm. All letters are uppercase so we have 26
// letters therefore a vector with a constant length of 26 is a good approach here.

class Solution {
public:
    static int characterReplacement(string s, int k) {
        int longest = 0;
        vector<int> chars(26, 0);
        int maxUnchangedNums = 1;

        for (int l = 0, r = 0; r < s.length(); r++) {
            chars[s[r] - 'A']++;

            int len = r + 1 - l;
            maxUnchangedNums = max(chars[s[r] - 'A'], maxUnchangedNums);
            int changesLeft = len - maxUnchangedNums;

            if (changesLeft > k) {
                chars[s[l] - 'A']--;
                l++;
            } else {
                longest++;
            }
        }

        return longest;
    }
};

int main() {
//    cout << Solution::characterReplacement("ABAB", 2) << "\n";
//    cout << Solution::characterReplacement("AABABBA", 1) << "\n";
//    cout << Solution::characterReplacement("BAAAB", 2) << "\n";

    string s = "AAAA";
    int count = 0;

    for (int i = 0; i < s.length(); i++) {
        if (s[i] == 'A') {
            count++;
        }
    }
    cout << count;
}