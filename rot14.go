package piscine

func runerot(x rune) rune {
	if x >= 'A' && x < 'M' || x >= 'a' && x < 'm' {
		return x + 14
	}
	if x >= 'M' && x <= 'Z' || x >= 'm' && x <= 'z' {
		return x - 12
	}
	return x
}

func Rot14(s string) string {
	resultat := ""
	for _, r := range s {
		resultat += string(runerot(r))
	}
	return resultat
}
