// https://leetcode.com/problems/merge-two-sorted-lists/?envType=study-plan&id=level-1
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
  ListNode *mergeTwoLists(ListNode *list1, ListNode *list2)
  {
    ListNode *dummyNode = new ListNode(INT_MIN);
    ListNode *next = dummyNode;

    while (list1 and list2)
    {
      if (list1->val <= list2->val)
      {
        next->next = list1;
        list1 = list1->next;
      }
      else
      {
        next->next = list2;
        list2 = list2->next;
      }
      next = next->next;
    }

    next->next = list1 ? list1 : list2;

    return dummyNode->next;
  }
};