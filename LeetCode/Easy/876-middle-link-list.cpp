#include <iostream>
#include <unordered_map>
#include <cmath>

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
    // BRUTE FORCE
    // space = O(n)
    // time = O(n)
    ListNode *middleNode(ListNode *head)
    {
        int count = 0;
        std::unordered_map<int, ListNode *> s;

        while (head != nullptr)
        {
            count++;
            s[count] = head;
            head = head->next;
        }

        return s[std::round(count / 2) + 1]; // we always want the upper half. Ints will be truncated
    }

    // FAST AND SLOW POINTER
    // space = O(1)
    // time = O(n)
    ListNode *middleNode2(ListNode *head)
    {
        ListNode *slow = head;
        ListNode *fast = head;

        while (fast != nullptr && fast->next != nullptr)
        {
            slow = slow->next;
            fast = fast->next->next;
        }
        return slow;
    }
};

int main()
{
    ListNode ListNode1 = ListNode(1);
    ListNode ListNode2 = ListNode(2);
    ListNode ListNode3 = ListNode(3);
    ListNode ListNode4 = ListNode(4);
    ListNode ListNode5 = ListNode(5);

    ListNode1.next = &ListNode2;
    ListNode2.next = &ListNode3;
    ListNode3.next = &ListNode4;
    ListNode4.next = &ListNode5;

    std::cout << Solution().middleNode(&ListNode1)->val;
}