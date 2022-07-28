// https://leetcode.com/problems/linked-list-cycle/
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

    // FLOYD's CYCLE ALGORITHM
    /**
     * Imagine two runners running on a track at different speed. What happens when the track is actually a circle?
     * The fast runner will eventually catch up to the slow runner
     *
     * space complexity = O(1) we only use two nodes
     * time complexity is O(n) if not cycle then we only execute n pointers otherwise
     * the fast pointer will have to do some extra iterations to catch up and lop the slow pointer (k) so time is
     * n + k so we can just say O(n)
     */
    bool hasCycle3(ListNode *head)
    {
        if (head == nullptr)
            return false;

        ListNode *slow = head;
        ListNode *fast = head->next;

        while (slow != fast)
        {
            if (fast == nullptr || fast->next == nullptr)
                return false;

            slow = slow->next;
            fast = fast->next->next;
        }
        return true;
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