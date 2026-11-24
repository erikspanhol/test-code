// Package textutil has small string helpers.
package textutil

// Truncate shortens s to at most n runes (not bytes). If s has more than n
// runes it is cut and a single "…" is appended; the ellipsis counts toward
// n, so the result never exceeds n runes. If n <= 0 it returns "".
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
