// https://leetcode.com/problems/add-two-numbers/
#include <iostream>

using namespace std;

struct ListNode
{
  int val;
  ListNode *next;
  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode *next) : val(x), next(next) {}
};

ListNode *addTwoNumbers(ListNode *l1, ListNode *l2)
{
  ListNode *startNode = new ListNode();
  ListNode *nextNode = startNode;
  while (nextNode != nullptr)
  {
    int l1Val = l1 ? l1->val : 0;
    int l2Val = l2 ? l2->val : 0;

    int sum = l1Val + l2Val + nextNode->val;
    int remainder = sum % (10);
    nextNode->val = remainder;
    int carryToNext = sum - remainder;

    if ((l1 and l1->next != nullptr) or (l2 and l2->next != nullptr) or (carryToNext) >= 10)
    {
      ListNode *next = new ListNode(carryToNext / 10);
      nextNode->next = next;
    }
    nextNode = nextNode->next;

    if (l1)
      l1 = l1->next;
    if (l2)
      l2 = l2->next;
  }

  return startNode;
}
