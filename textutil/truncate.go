// Package textutil has small string helpers.
package textutil

// Truncate shortens s to at most n characters, adding "…" when it cuts.
func Truncate(s string, n int) string {
	r := []rune(s)
	if n <= 0 {
		return ""
	}
	if len(r) <= n {
		return s
	}
	if n == 1 {
		return "…"
	}
	return string(r[:n-1]) + "…"
}
