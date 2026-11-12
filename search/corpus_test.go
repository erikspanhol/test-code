package search

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

// loadCorpus indexes search/testdata/corpus.txt (one "id<TAB>text" per line).
func loadCorpus(tb testing.TB) *Index {
	tb.Helper()
	f, err := os.Open("testdata/corpus.txt")
	if err != nil {
		tb.Fatal(err)
	}
	defer f.Close()
	ix := NewIndex()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		id, text, _ := strings.Cut(sc.Text(), "\t")
		ix.Add(Document{ID: id, Body: text})
	}
	return ix
}

func BenchmarkSearchCorpus(b *testing.B) {
	ix := loadCorpus(b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ix.Search("orchard cider")
	}
}
