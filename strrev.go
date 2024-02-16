package piscine

func StrRev(s string) string {
	strarray := []rune(s)
	lang := len(s)

	tap := ""
	for x := lang - 1; x >= 0; x-- {
		tap = string(strarray[x])
	}
	return tap
}
