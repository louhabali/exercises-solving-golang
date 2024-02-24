package piscine

func BasicAtoi(s string) int {
	arr := []rune(s)
	num := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] >= '0' && arr[i] <= '9' {
			num = num*10 + int(arr[i]-'0')
		}
	}
	return num
}
