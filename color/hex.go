package color

import "fmt"

func FromHex(s string) Color {
	if len(s) == 0 {
		return Auto
	}
	if s[0] == '#' {
		s = s[1:]
	}
	//func Sscanf(str string, format string, a ...interface{}) (n int, err error) {
	var r, g, b uint8
	n, _ := fmt.Sscanf(s, "%02x%02x%02x", &r, &g, &b)
	if n == 3 {
		return RGB(r, g, b)
	}
	return Auto
}
