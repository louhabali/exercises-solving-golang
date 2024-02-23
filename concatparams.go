package piscine

func ConcatParams(args []string) string {
	strarray := make([]byte, 0)
	for index, arg := range args {
		strarray = append(strarray, arg...)
		if index < len(args)-1 {
			strarray = append(strarray, '\n')
		}

	}
	return string(strarray)
}
