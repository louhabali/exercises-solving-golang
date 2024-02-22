package piscine

func TrimAtoi(s string) int {
	arr := []rune(s)
	var num int = 0
	signe := 1
	for i := 0; i < len(arr); i++ {
		if arr[i] >= '0' && arr[i] <= '9' {
			D := int(arr[i] - '0')
			num = num*10 + D
		}
		if arr[i] == '-' && num == 0 {
			signe = -1
		}
	}
	return signe * num
}
