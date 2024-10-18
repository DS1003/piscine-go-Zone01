package piscine

type NodeL6 struct {
	Data interface{}
	Next *NodeL6
}

func ListAt(l *NodeL, pos int) *NodeL {
	n := 0
	for l != nil {
		if n == pos {
			return l
		}
		n++
		l = l.Next
	}
	return nil
}
