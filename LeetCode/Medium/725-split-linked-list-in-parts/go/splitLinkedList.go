package splitLinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func SplitListToParts(head *ListNode, k int) []*ListNode {
	listLen := 0
	tempHead := head

	for tempHead != nil {
		listLen++
		tempHead = tempHead.Next
	}

	res := make([]*ListNode, k)
	segmentLen := listLen / k
	remainder := listLen % k
	pos := 0

	for head != nil {
		res[pos] = head
		var last *ListNode = nil
		groupLen := segmentLen
		if pos < remainder {
			groupLen += 1
		}

		for i := 0; i < groupLen; i++ {
			last = head
			head = head.Next
		}
		if last != nil {
			last.Next = nil
		}
		pos++
	}

	return res
}
