package splitLinkedList_test

import (
	splitLinkedList "github.com/rdforte/competitive-programming/LeetCode/Medium/725-split-linked-list-in-parts/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitLinkedListIntoPartsShortList(t *testing.T) {
	node3 := splitLinkedList.ListNode{Val: 3}
	node2 := splitLinkedList.ListNode{Val: 2, Next: &node3}
	head := splitLinkedList.ListNode{Val: 1, Next: &node2}

	numParts := 5
	output := splitLinkedList.SplitListToParts(&head, numParts)
	expected := []*splitLinkedList.ListNode{
		{1, nil},
		{2, nil},
		{3, nil},
	}

	assert.Equal(t, numParts, len(output))

	for i := 0; i < len(expected); i++ {
		assert.Equal(t, expected[i].Val, output[i].Val)
		assert.Equal(t, expected[i].Next, output[i].Next)
	}
}

func TestSplitLinkedListIntoPartsLongList(t *testing.T) {
	node10 := splitLinkedList.ListNode{Val: 10}
	node9 := splitLinkedList.ListNode{Val: 9, Next: &node10}
	node8 := splitLinkedList.ListNode{Val: 8, Next: &node9}
	node7 := splitLinkedList.ListNode{Val: 7, Next: &node8}
	node6 := splitLinkedList.ListNode{Val: 6, Next: &node7}
	node5 := splitLinkedList.ListNode{Val: 5, Next: &node6}
	node4 := splitLinkedList.ListNode{Val: 4, Next: &node5}
	node3 := splitLinkedList.ListNode{Val: 3, Next: &node4}
	node2 := splitLinkedList.ListNode{Val: 2, Next: &node3}
	head := splitLinkedList.ListNode{Val: 1, Next: &node2}

	numParts := 3
	output := splitLinkedList.SplitListToParts(&head, numParts)

	expectNode10 := splitLinkedList.ListNode{Val: 10}
	expectNode9 := splitLinkedList.ListNode{Val: 9, Next: &expectNode10}
	expectNode8 := splitLinkedList.ListNode{Val: 8, Next: &expectNode9}
	expectNode7 := splitLinkedList.ListNode{Val: 7}
	expectNode6 := splitLinkedList.ListNode{Val: 6, Next: &expectNode7}
	expectNode5 := splitLinkedList.ListNode{Val: 5, Next: &expectNode6}
	expectNode4 := splitLinkedList.ListNode{Val: 4}
	expectNode3 := splitLinkedList.ListNode{Val: 3, Next: &expectNode4}
	expectNode2 := splitLinkedList.ListNode{Val: 2, Next: &expectNode3}
	expectHead := splitLinkedList.ListNode{Val: 1, Next: &expectNode2}
	expected := []*splitLinkedList.ListNode{
		&expectHead,
		&expectNode5,
		&expectNode8,
	}

	assert.Equal(t, numParts, len(output))

	for i := 0; i < len(expected); i++ {
		assert.Equal(t, expected[i].Val, output[i].Val)
		assert.Equal(t, expected[i].Next, output[i].Next)
	}
}
