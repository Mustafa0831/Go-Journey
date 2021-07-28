package student

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	l.Tail.Next = n
	l.Tail = n

}

func ListRemoveIf(l *List, data_ref interface{}) {
	mover := l.Head
	correct := l.Head
	for mover != nil {
		Next := mover.Next
		if Next == nil {
			correct.Next = nil
			break
		}
		if Next.Data != data_ref {
			correct.Next = Next
			correct = correct.Next
		}
		mover = mover.Next
	}
}
