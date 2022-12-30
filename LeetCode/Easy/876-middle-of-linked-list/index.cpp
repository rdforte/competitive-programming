#include <bits/stdc++.h>

using namespace std;

struct ListNode
{
  int val;
  ListNode *next;
  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode *next) : val(x), next(next) {}
};

// Space is O(1)
// time is O(n) but involves doing O(n) + O(0.5*n)
class SolutionOne
{
public:
  ListNode *middleNode(ListNode *head)
  {
    int count = 0;
    ListNode *counterHead = head;

    while (counterHead)
    {
      count++;
      counterHead = counterHead->next;
    }

    count /= 2;
    while (count)
    {
      count--;
      head = head->next;
    }

    return head;
  }
};

// most optimal as time is O(n) but can do O(0.5*n) as fast pointer moves 2 places.
class Solution
{
public:
  ListNode *middleNode(ListNode *head)
  {
    ListNode *fast = head;
    ListNode *slow = head;

    while (fast and fast->next)
    {
      fast = fast->next->next;
      slow = slow->next;
    }
    return slow;
  }
};