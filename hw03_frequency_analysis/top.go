package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"regexp"
	"sort"
	"strings"
)

func Top10(words string) []string {
	var result []string
	if len(words) == 0 {
		return result
	}

	r := regexp.MustCompile("\\s*\\d+\\s*|[\\s*\\W*\\s*]+")
	words = r.ReplaceAllString(words, " ")
	if strings.TrimSpace(words) == "" {
		return result
	}

	wordsSlice := strings.Split(string(words), " ")
	wordsCounted := make(map[string]int)
	for _, word := range wordsSlice {
		wordsCounted[strings.ToLower(word)]++
	}

	TopWords := []struct {
		Word  string
		Count int
	}{}

	for word, count := range wordsCounted {
		TopWords = append(TopWords, struct {
			Word  string
			Count int
		}{Word: word, Count: count})
	}

	sort.Slice(TopWords, func(i, j int) bool {
		return TopWords[i].Count > TopWords[j].Count
	})

	countOfTop := len(TopWords)
	if len(TopWords) > 10 {
		countOfTop = 10
	}

	result = make([]string, 10)
	for i := 0; i < countOfTop; i++ {
		result[i] = TopWords[i].Word
	}

	return result
}
