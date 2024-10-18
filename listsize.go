package piscine

type NodeL3 struct {
	Data interface{}
	Next *NodeL
}

type List3 struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	n := 0
	for l.Head != nil {
		n++
		l.Head = l.Head.Next
	}
	return n
}
