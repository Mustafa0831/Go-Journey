package student

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	var n3 *NodeI

	for n1 != nil {
		n3 = SortListInsert(n3, n1.Data)
		n1 = n1.Next
	}
	for n2 != nil {
		n3 = SortListInsert(n3, n2.Data)
		n2 = n2.Next
	}
	return n3
}
