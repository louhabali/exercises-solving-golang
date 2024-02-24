package piscine

func Atoi(s string) int {
	arr := []rune(s)
	num := 0
	sg := 1
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] >= '0' && arr[i] <= '9' {
			num = num*10 + int(arr[i]-'0')
		}
		if num == 0 && arr[i] == '-' {
			sg = -1
			count++
		}
		if num == 0 && arr[i] == '+' {
			count++
		}
		if !(arr[i] >= '0' && arr[i] <= '9') && (count > 1) {
			return 0
		}

	}

	return sg * num
}
