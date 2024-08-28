package trie

import "testing"

func TestInsertSearch(t *testing.T) {
	trie := Constructor()
	word := "test"
	trie.Insert(word)
	if !trie.Search(word) {
		t.Fatalf("expected %q to exist in trie", word)
	}
}
func TestStartsWith(t *testing.T) {
	trie := Constructor()
	word := "test"
	trie.Insert(word)

	prefix := "te"
	if !trie.StartsWith(prefix) {
		t.Fatalf("expected trie to start with %q", prefix)
	}
}
