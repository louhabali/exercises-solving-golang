package piscine

func ToLower(s string) string {
	arr := []rune(s)
	for i := 0; i <= len(s)-1; i++ {
		if arr[i] >= 'A' && arr[i] <= 'Z' {
			arr[i] += 32
		}
	}
	return string(arr)
}
