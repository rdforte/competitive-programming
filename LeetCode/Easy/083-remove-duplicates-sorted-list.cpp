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
    // runtime = O(n) we just have to run through all the nodes once removing the ones in front if they are the same.
    // once removed then we skip over all the ones removed and go straight to the node with a different value.
    // space complexity = O(1) we just have to keep track of the cur (current node)
    ListNode *deleteDuplicates(ListNode *head)
    {
        ListNode *cur = head;

        while (cur != nullptr && cur->next != nullptr)
        {
            if (cur->val == cur->next->val)
            {
                cur->next = cur->next->next;
            }
            else
            {
                cur = cur->next;
            }
        }
        return head;
    }
};

int main()
{
    ListNode ListNode1 = ListNode(1);
    ListNode ListNode2 = ListNode(2);
    ListNode ListNode3 = ListNode(2);
    ListNode ListNode4 = ListNode(2);
    ListNode ListNode5 = ListNode(3);

    ListNode1.next = &ListNode2;
    ListNode2.next = &ListNode3;
    ListNode3.next = &ListNode4;
    ListNode4.next = &ListNode5;

    ListNode *n = Solution().deleteDuplicates(&ListNode1);

    while (n != nullptr)
    {
        std::cout << n->val << " ";
        n = n->next;
    }
}