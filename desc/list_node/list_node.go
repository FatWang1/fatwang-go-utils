package list_node

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListNodeify(list ...int) *ListNode {
	if len(list) == 0 {
		return nil
	}
	head := &ListNode{
		Val: list[0],
	}
	curr := head
	for i := 1; i < len(list); i++ {
		curr.Next = &ListNode{
			Val: list[i],
		}
		curr = curr.Next
	}
	return head
}
