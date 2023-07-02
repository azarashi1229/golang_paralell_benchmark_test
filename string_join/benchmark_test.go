package random

import (
	"strconv"
	"strings"
	"testing"
)

var args = func() []string {
	r := make([]string, 100)
	for i := range r {
		r[i] = strconv.Itoa(i)
	}
	return r
}()

func BenchmarkNotUseStringJoin(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		notUseJoin(args)
	}
}
func BenchmarkUseStringJoin(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		useJoin(args)
	}
}

func useJoin(args []string) string {
	return strings.Join(args, " ")
}
func notUseJoin(args []string) string {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	return s
}
