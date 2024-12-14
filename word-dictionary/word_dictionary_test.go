package word_dictionary

import "testing"

func TestWordDictionary(t *testing.T) {
	wordDictionary := Constructor()

	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")

	if wordDictionary.Search("pad") {
		t.Errorf("did not expect %q", "pad")
	}

	if !wordDictionary.Search("bad") {
		t.Errorf("expected %q", "bad")
	}

	if !wordDictionary.Search(".ad") {
		t.Errorf("expected %q", ".ad")
	}

	if !wordDictionary.Search("b..") {
		t.Errorf("expected %q", "b..")
	}
}
