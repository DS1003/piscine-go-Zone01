package piscine

func ConcatParams(args []string) string {
	var str string
	for i, j := range args {
		str += string(j)
		if i != len(args)-1 {
			str += "\n"
		}
	}
	return str
}
