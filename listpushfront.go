package piscine

type Node struct {
	Data interface{}
	Next *NodeL
}

type List1 struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	newnode := &NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = newnode
	} else {
		newnode.Next = l.Head
		l.Head = newnode

	}
}
