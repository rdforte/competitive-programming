#include <iostream>
#include <vector>
#include <set>

struct ListNode
{
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL){}; // member initializer list using uniform initialization
};

class Solution
{
public:
    // BRUTE FORCE APPROACH
    // for every loop of n we have to do another loop of n therefore time complexity is n^2
    // space time is n because of the vector
    bool hasCycle(ListNode *head)
    {
        if (head == nullptr || head->next == nullptr)
            return false;

        std::vector<ListNode *> ln = {};

        while (head->next != NULL)
        {
            for (ListNode *n : ln)
            {
                if (head == n)
                    return true;
            }

            ln.push_back(head);
            head = head->next;
        }
        return false;
    }

    // We use a set ie: hash map
    // runtime is n because we only loop through n once
    // spacetime is n because we have to keep track of each node in a set
    bool hasCycle2(ListNode *head)
    {
        if (head == nullptr || head->next == nullptr)
            return false;

        std::set<ListNode *> s{};

        while (head->next != NULL)
        {

            if (s.find(head) != s.end())
            {
                return true;
            }

            s.insert(head);
            head = head->next;
        }
        return false;
    }
};

int main()
{
    ListNode ListNode1 = ListNode(1);
    ListNode ListNode2 = ListNode(2);
    ListNode ListNode3 = ListNode(3);
    ListNode ListNode4 = ListNode(4);

    ListNode1.next = &ListNode2;
    ListNode2.next = &ListNode3;
    ListNode3.next = &ListNode4;
    // ListNode4.next = &ListNode1;

    std::cout
        << Solution().hasCycle2(&ListNode1);
}