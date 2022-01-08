#include <iostream>

struct ListNode
{
    int val;
    ListNode *next;

    ListNode() : val(0), next(nullptr){};
    ListNode(int x) : val(x), next(nullptr){};
    ListNode(int x, ListNode *next) : val(x), next(next){};
};

class Solution
{
public:
    // runTime = O(n) we have to iterate through every node
    // spaceTime = O(1) we only have to keep track of h
    ListNode *reverseList(ListNode *head)
    {
        ListNode *h = nullptr;

        while (head != nullptr)
        {
            ListNode *t = head->next;
            head->next = h;
            h = head;
            head = t;
        }
        return h;
    }
};

int main()
{
    ListNode ListNode1 = ListNode(1);
    ListNode ListNode2 = ListNode(2);
    // ListNode ListNode3 = ListNode(3);

    ListNode1.next = &ListNode2;
    // ListNode2.next = &ListNode3;

    ListNode *n = Solution().reverseList(&ListNode1);

    // ListNode *n = &ListNode1;

    while (n != nullptr)
    {
        std::cout << n->val << " ";
        n = n->next;
    }
}