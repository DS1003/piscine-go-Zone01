package piscine

func Compact(ptr *[]string) int {
	r := *ptr
	count := 0
	for _, i := range *ptr {
		if i != "" {
			r[count] = i
			count++
		}
	}
	*ptr = r[0:count]
	return count
}
