package piscine

func ForEach(f func(int), arg []int) {
	for _, r := range arg {
		f(r)
	}
}
