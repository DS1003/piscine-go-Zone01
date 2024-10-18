package piscine

type NodeL10 struct {
	Data interface{}
	Next *NodeL
}

type List10 struct {
	Head *NodeL
	Tail *NodeL
}

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	s := l.Head
	for s != nil {
		if comp(s.Data, ref) {
			return &s.Data
		}
		s = s.Next
	}
	return nil
}
