package piscine

func BasicAtoi2(s string) int {
	arr := []rune(s)
	num := 0
	for i := 0; i < len(arr); i++ {
		if !(arr[i] >= '0' && arr[i] <= '9') {
			return 0
		}
		if arr[i] >= '0' && arr[i] <= '9' {
			num = num*10 + int(arr[i]-'0')
		}
	}
	return num
}
