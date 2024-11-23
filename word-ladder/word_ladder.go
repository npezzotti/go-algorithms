package word_ladder

type Word struct {
	word  string
	level int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]struct{}, len(wordList))
	for _, word := range wordList {
		wordSet[word] = struct{}{}
	}

	queue := make([]Word, 0)
	queue = append(queue, Word{word: beginWord, level: 1})

	visited := make(map[string]bool, len(wordList))
	visited[beginWord] = true

	for len(queue) > 0 {
		word := queue[0]
		queue = queue[1:]

		if word.word == endWord {
			return word.level
		}

		charList := []rune(word.word)
		for i := range charList {
			original := charList[i]

			for c := 'a'; c <= 'z'; c++ {
				charList[i] = c
				newWord := string(charList)

				_, ok := wordSet[newWord]
				if ok && !visited[newWord] {
					visited[newWord] = true
					queue = append(queue, Word{word: newWord, level: word.level + 1})
				}
			}

			charList[i] = original
		}
	}

	return 0
}
