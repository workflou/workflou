package color

import "strings"

var letterHues = map[string]int{
	"a": 14, "b": 197, "c": 271, "d": 45, "e": 225,
	"f": 154, "g": 84, "h": 314, "i": 174, "j": 291,
	"k": 111, "l": 240, "m": 61, "n": 340, "o": 28,
	"p": 197, "q": 126, "r": 0, "s": 170, "t": 280,
	"u": 77, "v": 220, "w": 94, "x": 317, "y": 140,
	"z": 260,
}

func HueForInitial(s string) int {
	if len(s) == 0 {
		return 200
	}

	firstLetter := strings.ToLower(string([]rune(s)[0]))
	if hue, ok := letterHues[firstLetter]; ok {
		return hue
	}

	return 200
}
