// Package search is a small in-memory document index.
package search

import (
	"sort"
	"strings"
)

// Document is something that can be searched.
type Document struct {
	ID    string
	Title string
	Body  string
}

// Index holds documents and answers keyword queries.
type Index struct {
	docs []Document
}

// NewIndex returns an empty index.
func NewIndex() *Index {
	return &Index{}
}

// Add stores a document. Adding a document with an existing ID replaces it.
func (ix *Index) Add(d Document) {
	for i := range ix.docs {
		if ix.docs[i].ID == d.ID {
			ix.docs[i] = d
			return
		}
	}
	ix.docs = append(ix.docs, d)
}

// Len returns the number of documents.
func (ix *Index) Len() int {
	return len(ix.docs)
}

// Search returns the IDs of documents containing every word of the query,
// best match first. Matching is case-insensitive.
func (ix *Index) Search(query string) []string {
	words := strings.Fields(strings.ToLower(query))
	if len(words) == 0 {
		return nil
	}
	type hit struct {
		id    string
		score int
	}
	var hits []hit
	for _, d := range ix.docs {
		tokens := tokenize(d.Title + " " + d.Body)
		score := 0
		matchedAll := true
		for _, w := range words {
			n := tokens[w]
			if n == 0 {
				matchedAll = false
				break
			}
			score += n
		}
		if matchedAll {
			hits = append(hits, hit{d.ID, score})
		}
	}
	sort.Slice(hits, func(i, j int) bool {
		if hits[i].score != hits[j].score {
			return hits[i].score > hits[j].score
		}
		return hits[i].id < hits[j].id
	})
	ids := make([]string, len(hits))
	for i, h := range hits {
		ids[i] = h.id
	}
	return ids
}

// tokenize lowercases text and counts each word.
func tokenize(text string) map[string]int {
	counts := map[string]int{}
	for _, f := range strings.FieldsFunc(strings.ToLower(text), isSeparator) {
		counts[f]++
	}
	return counts
}

func isSeparator(r rune) bool {
	return !(r >= 'a' && r <= 'z' || r >= '0' && r <= '9')
}
