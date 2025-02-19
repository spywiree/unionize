package parse

import (
	"unicode/utf8"

	"github.com/spywiree/unionize/unionize/internal"
)

func mapFirstRune(s string, fn func(rune) rune) string {
	r, size := utf8.DecodeRuneInString(s)
	buf := utf8.AppendRune(nil, fn(r))
	buf = append(buf, s[size:]...)
	return internal.BytesToString(buf)
}
