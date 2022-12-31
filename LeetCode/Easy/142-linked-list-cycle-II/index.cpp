#include <bits/stdc++.h>

using namespace std;

struct ListNode
{
  int val;
  ListNode *next;
  ListNode(int x) : val(x), next(NULL) {}
};

// Brute force using set. time is O(n) and space O(n)
class SolutionBruteForce
{
public:
  ListNode *detectCycle(ListNode *head)
  {
    set<ListNode *> visited;
    while (head)
    {
      if (visited.find(head) != visited.end())
      {
        return head;
      }
      visited.insert(head);
      head = head->next;
    }
    return NULL;
  }
};

// Optimal Solution uses O(n) runtime and O(1) space.
class Solution
{
public:
  ListNode *detectCycle(ListNode *head)
  {
    ListNode *cyclicNode = findCyclicNode(head);
    if (!cyclicNode)
      return NULL;

    while (head != cyclicNode)
    {
      head = head->next;
      cyclicNode = cyclicNode->next;
    }

    return cyclicNode;
  }

  ListNode *findCyclicNode(ListNode *head)
  {
    ListNode *fast = head;
    ListNode *slow = head;
    while (fast and fast->next)
    {
      fast = fast->next->next;
      slow = slow->next;
      if (fast == slow)
      {
        return fast;
      }
    }

    return NULL;
  }
};

int main()
{
  ListNode *one = new ListNode(3);
  ListNode *two = new ListNode(2);
  ListNode *three = new ListNode(0);
  ListNode *four = new ListNode(-4);

  one->next = two;
  two->next = three;
  three->next = four;
  four->next = two;

  cout << Solution().detectCycle(one);
}