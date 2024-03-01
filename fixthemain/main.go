package piscine

import "github.com/01-edu/z01"

type Door struct {
	State string
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...")
	ptrDoor.State = "OPEN"
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	ptrDoor.State = "CLOSED"
}

func IsDoorOpen(door Door) bool {
	PrintStr("Is the Door opened ?")
	return door.State == "OPEN"
}

func IsDoorClose(door Door) bool {
	PrintStr("Is the Door closed ?")
	return door.State == "CLOSED"
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(*door) {
		OpenDoor(door)
	}
	if IsDoorOpen(*door) {
		CloseDoor(door)
	}
	if door.State == "OPEN" {
		CloseDoor(door)
	}
}
