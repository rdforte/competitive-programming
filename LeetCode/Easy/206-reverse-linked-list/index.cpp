#include "../../../stdc++.h"

using namespace std;

struct ListNode
{
  int val;
  ListNode *next;

  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode *next) : val(x), next(next) {}
};

ListNode *reverseList(ListNode *head)
{
  ListNode *prevHead = nullptr;
  while (head != nullptr)
  {
    ListNode *next = head->next;
    head->next = prevHead;
    prevHead = head;
    head = next;
  }
  return prevHead;
}

int main()
{
}