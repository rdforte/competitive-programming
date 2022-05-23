// https://leetcode.com/problems/merge-two-sorted-lists/
#include <iostream>

struct ListNode
{
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr){};
    ListNode(int x) : val(x), next(nullptr){};
    ListNode(int x, ListNode *next) : val(0), next(next){};
};

class Solution
{
public:
    // space complexity is keeping track of start and prev so O(n).
    // time complexity is the combination of the two loops so list1(n) && list2(m) gives us O(n + m)
    ListNode *mergeTwoLists(ListNode *list1, ListNode *list2)
    {
        ListNode *start = new ListNode(-1);
        ListNode *prev = start;

        while (list1 != nullptr && list2 != nullptr)
        {
            if (list1->val <= list2->val)
            {
                prev->next = list1;
                list1 = list1->next;
            }
            else
            {
                prev->next = list2;
                list2 = list2->next;
            }
            prev = prev->next;
        }

        // we might still have some remaining nodes
        prev->next = list1 == NULL ? list2 : list1;

        return start->next;
    }
};

int main()
{
    ListNode list1 = ListNode(1);
    ListNode list2 = ListNode(2);

    ListNode *n = Solution().mergeTwoLists(&list1, &list2);

    std::cout << "---------------------\n";

    while (n != nullptr)
    {
        std::cout << n->val << " ";
        n = n->next;
    }
}