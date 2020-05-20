package linkedlists

type LinkedList struct {
	Head *node
}

type node struct {
	data int
	next *node
}

func (l *LinkedList) Append(data int) {
	if l.Head == nil {
		l.Head = &node{data: data, next: nil}
		return
	}

	cur := l.Head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = &node{data: data, next: nil}
}

func (l *LinkedList) Prepend(data int) {
	newHead := &node{
		data: data,
		next: l.Head,
	}

	l.Head = newHead
}

func (l *LinkedList) DeleteWithValue(data int) {
	if l.Head == nil {
		return
	}

	if l.Head.data == data {
		l.Head = l.Head.next
		return
	}

	cur := l.Head
	for cur.next != nil {
		if cur.next.data == data {
			cur.next = cur.next.next
			return
		}

		cur = cur.next
	}
}
