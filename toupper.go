package piscine

func ToUpper(s string) string {
	arr := []rune(s)
	for i := 0; i <= len(s)-1; i++ {
		if arr[i] >= 'a' && arr[i] <= 'z' {
			arr[i] -= 32
		}
	}
	return string(arr)
}
