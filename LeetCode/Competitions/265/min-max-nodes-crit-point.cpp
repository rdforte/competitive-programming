#include <iostream>
#include <vector>

class Node
{
public:
    int val;
    Node *next;
    Node(int v, Node *n)
    {
        val = v;
        next = n;
    }
};

class Solution
{
public:
    std::vector<int> nodesBetweenCriticalPoints(Node *head)
    {

        for (int i = 1; head != nullptr; i++)
        {
            std::cout << head->val << " ";
            head = head->next;
        }
        return std::vector<int>{1};
    }
};

int main()
{
    Node n3 = Node(3, nullptr);
    Node n2 = Node(2, &n3);
    Node n1 = Node(1, &n2);

    Solution().nodesBetweenCriticalPoints(&n1);
}
