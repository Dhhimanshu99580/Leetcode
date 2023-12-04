func countCharacters(words []string, chars string) int {
    charCount := make(map[rune]int)
	for _, char := range chars {
		charCount[char]++
	}
	canFormWord := func(word string) bool {
		wordCount := make(map[rune]int)
		for _, char := range word {
			wordCount[char]++
		}
		for char, count := range wordCount {
			if charCount[char] < count {
				return false
			}
		}

		return true
	}
	sumLength := 0
	for _, word := range words {
		if canFormWord(word) {
			sumLength += len(word)
		}
	}

	return sumLength
}
