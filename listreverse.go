package piscine

type NodeL7 struct {
	Data interface{}
	Next *NodeL
}

type List7 struct {
	Head *NodeL
	Tail *NodeL
}

func ListReverse(l *List) {
	newprtr := l.Head
	s := l.Head
	s = nil
	for newprtr != nil {
		n := newprtr.Next
		newprtr.Next = s
		s = newprtr
		newprtr = n
	}
	l.Head = l.Tail
	l.Tail = s
}
