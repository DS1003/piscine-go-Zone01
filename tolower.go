package piscine

func ToLower(s string) string {
	runelow := []rune(s)
	for i := range runelow {
		if runelow[i] >= 'A' && runelow[i] <= 'Z' {
			runelow[i] = runelow[i] + 32
		}
	}
	return string(runelow)
}
