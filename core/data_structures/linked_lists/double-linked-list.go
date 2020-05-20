package linkedlists

type DoubleLinkedList struct {
	Head *dnode
}

type dnode struct {
	data int
	prev *dnode
	next *dnode
}

func (l *DoubleLinkedList) Append(data int) {
	l.Head = nil
}
