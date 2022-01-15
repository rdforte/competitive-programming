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
    ListNode *removeElements(ListNode *head, int val)
    {
        ListNode *prev = NULL;
        ListNode *start = head;

        while (head != nullptr)
        {
            if (head->val == val)
            {
                if (prev == NULL)
                {
                    start = head->next;
                    head = start;
                }
                else
                {
                    prev->next = head->next;
                    head = head->next;
                }
            }
            else
            {
                prev = head;
                head = head->next;
            }
        }
        return start;
    };
};

int main()
{
    ListNode ListNode1 = ListNode(1);
    ListNode ListNode2 = ListNode(2);
    ListNode ListNode3 = ListNode(1);
    ListNode ListNode4 = ListNode(2);
    ListNode ListNode5 = ListNode(1);

    ListNode1.next = &ListNode2;
    ListNode2.next = &ListNode3;
    ListNode3.next = &ListNode4;
    ListNode4.next = &ListNode5;

    ListNode *n = Solution().removeElements(&ListNode1, 1);

    while (n != nullptr)
    {
        std::cout << n->val << " ";
        n = n->next;
    }
}