package main

import (
	"strings"
	"unicode"
)

var stopWords = map[string]struct{}{
	"a": {}, "and": {}, "am": {}, "be": {}, "have": {}, "i": {},
	"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
}

func analyze(text string) []string {
	tokens := tokenizer(text)
	tokens = toLowerCase(tokens)
	tokens = stopWordsFilter(tokens)
	return tokens
}

// tokenizer
// hey how are you?
// [hey, how, are, you]
func tokenizer(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

//******************
// 		filters
//******************

//toLowerCase
func toLowerCase(tokens []string) []string {
	r := make([]string, len(tokens))

	for i, t := range tokens {
		r[i] = strings.ToLower(t)
	}
	return r
}

//stopWordsFilter
func stopWordsFilter(tokens []string) []string {
	r := make([]string, 0, len(tokens))

	for _, token := range tokens {
		if _, ok := stopWords[token]; !ok {
			r = append(r, token)
		}
	}
	return r
}
