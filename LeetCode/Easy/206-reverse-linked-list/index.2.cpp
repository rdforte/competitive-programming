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

class Solution
{
public:
  ListNode *reverseList(ListNode *head)
  {
    ListNode *prev = nullptr;
    ListNode *start = nullptr;

    while (head)
    {
      ListNode *temp = head;
      head = head->next;
      temp->next = prev;
      prev = temp;
      start = prev;
    }

    return start;
  }
};