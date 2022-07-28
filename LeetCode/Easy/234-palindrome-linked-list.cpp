// https://leetcode.com/problems/palindrome-linked-list/
#include <vector>
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
    // brute force
    // I have to do 2 loops so O(n) runtime
    // I have to keep track of the list of items so space complexity is O(n)
    bool isPalindrome(ListNode *head)
    {
        std::vector<int> v = {};

        while (head != nullptr)
        {
            v.push_back(head->val);
            head = head->next;
        };

        int revP = v.size() - 1;

        for (int i = 0; i < v.size(); i++)
        {
            if (v[i] != v[revP])
                return false;
            revP--;
        }

        return true;
    }

    // two runners pointer technique
    // spaceComplexity = O(1) we only have to keep track of the slow, fast and rev node.
    // timeComplexity = O(n)
    // This method saves on memory but it is not as performant as the first solution
    bool isPalindromeV2(ListNode *head)
    {

        if (head == nullptr)
            return true;

        ListNode *slow = head;
        ListNode *fast = head;

        // 1. This will put the slow pointer in the middle of the linked list
        while (fast != nullptr && fast->next != nullptr)
        {
            slow = slow->next;
            fast = fast->next->next;
        }

        // 2. reverse the second half of the list
        ListNode *rev = reverseList(slow);

        // 3. compare the start to the reversed end
        while (rev != nullptr)
        {
            if (rev->val != head->val)
                return false;

            head = head->next;
            rev = rev->next;
        }

        return true;
    }

private:
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
    ListNode ListNode3 = ListNode(3);
    ListNode ListNode4 = ListNode(2);
    ListNode ListNode5 = ListNode(1);

    ListNode1.next = &ListNode2;
    ListNode2.next = &ListNode3;
    ListNode3.next = &ListNode4;
    ListNode4.next = &ListNode5;

    // std::cout << Solution().isPalindrome(&ListNode1);
    Solution().isPalindromeV2(&ListNode1);
}