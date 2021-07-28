package student

type NodeI struct {
	Data int
	Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	node := &NodeI{Data: data_ref}
	if l == nil {
		return node
	}
	if node.Data < l.Data {
		node.Next = l
		return node
	}
	Head := l
	for l.Next != nil {
		if node.Data < l.Next.Data {
			node.Next = l.Next
			l.Next = node
			return Head
		}
		l = l.Next
	}
	l.Next = node
	return Head
}