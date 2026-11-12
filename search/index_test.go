package search

import (
	"reflect"
	"testing"
)

func sample() *Index {
	ix := NewIndex()
	ix.Add(Document{ID: "a", Title: "apple pie", Body: "a recipe with apple and cinnamon"})
	ix.Add(Document{ID: "b", Title: "banana bread", Body: "banana, flour and an apple"})
	ix.Add(Document{ID: "c", Title: "cherry tart", Body: "cherry and sugar"})
	return ix
}

func TestSearchSingleWord(t *testing.T) {
	got := sample().Search("apple")
	want := []string{"a", "b"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Search(apple) = %v, want %v", got, want)
	}
}

func TestSearchQueryIsCaseInsensitive(t *testing.T) {
	got := sample().Search("Apple")
	want := []string{"a", "b"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Search(Apple) = %v, want %v", got, want)
	}
}

func TestSearchAllWordsMustMatch(t *testing.T) {
	got := sample().Search("apple banana")
	want := []string{"b"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Search(apple banana) = %v, want %v", got, want)
	}
}

func TestSearchEmptyQuery(t *testing.T) {
	if got := sample().Search("   "); got != nil {
		t.Fatalf("empty query returned %v", got)
	}
}

func TestAddReplacesByID(t *testing.T) {
	ix := sample()
	ix.Add(Document{ID: "c", Title: "plum", Body: "plum"})
	if ix.Len() != 3 {
		t.Fatalf("Len = %d, want 3", ix.Len())
	}
	if got := ix.Search("cherry"); len(got) != 0 {
		t.Fatalf("replaced document still found: %v", got)
	}
}
