package piscine

type NodeL4 struct {
	Data interface{}
	Next *NodeL
}

type List4 struct {
	Head *NodeL
	Tail *NodeL
}

func ListLast(l *List) interface{} {
	if l.Head != nil {
		for l.Head.Next != nil {
			l.Head = l.Head.Next
		}
		return l.Head.Data
	}
	return nil
}
