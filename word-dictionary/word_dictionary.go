package word_dictionary

type WordDictionary struct {
	children [26]*WordDictionary
	isEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (wd *WordDictionary) AddWord(word string) {
	cur := wd
	for _, char := range word {
		idx := char - 'a'
		if cur.children[idx] == nil {
			cur.children[idx] = &WordDictionary{}
		}

		cur = cur.children[idx]
	}

	cur.isEnd = true
}

func (wd *WordDictionary) Search(word string) bool {
	cur := wd
	for i, char := range word {
		switch char {
		case '.':
			for _, child := range cur.children {
				if child != nil {
					if child.Search(word[i+1:]) {
						return true
					}
				}
			}
			return false
		default:
			idx := char - 'a'
			if cur.children[idx] == nil {
				return false
			}

			cur = cur.children[idx]
		}
	}

	return cur.isEnd
}
