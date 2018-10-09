package repeat

import (
	"strings"
)

// Repeat - Return a string with char repeated.
func Repeat(char string, count int) string {
	return strings.Repeat(char, count)
}
