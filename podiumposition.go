package piscine

func PodiumPosition(podium [][]string) [][]string {
	n := len(podium)
	for i := 0; i < n/2; i++ {
		podium[i], podium[n-i-1] = podium[n-i-1], podium[i]
	}
	return podium
}
