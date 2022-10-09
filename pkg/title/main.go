// Title capitalization package.
//
// Contains functions used for reformatting input text following English title
// capitalization rules. First word, last word, new sentences, punctuated
// abbreviations, brackets () {} [] <>, quotes " ', etc is handled correctly.
// Corner case proofness is very good but should NOT be expected, it is more or
// less impossible to achieve correct results on any input.
//
// Lowercase:
//
//	a, an, and, as, at, but, by, en, etc, for, from,
//	if, in, of, on, or, the, to, via, von, vs, with
//
// Example:
//
//	Input:  "tears for fears @ rule the world: the greatest hits"
//	Output: "Tears for Fears @ Rule the World: The Greatest Hits"
//
// Usage:
//
//	import "github.com/ggustafsson/godis/pkg/title"
//
//	title.Titleize(string)
//
// Author: Göran Gustafsson <gustafsson.g@gmail.com>
//
// License: BSD 3-Clause
package title

import (
	"regexp"
	"strings"

	"github.com/ggustafsson/godis/pkg/godis"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var lowercase = []string{
	"a", "an", "and", "as", "at", "but", "by", "en", "etc", "for", "from",
	"if", "in", "of", "on", "or", "the", "to", "via", "von", "vs", "with",
}
var end_sentence = []string{
	".", ",", ":", ";", "!", "?", "&", "/", "+", "-",
}

// Capitalize every word (including sub-words) in string. Same behaviour as
// cases.Title() + x.String() except it handles punctuated abbreviations.
func Title(word string) string {
	caser := cases.Title(language.English)
	// XXX: This feels dirty. Is there a better way?
	words := strings.Split(caser.String(word), ".")
	for index, word := range words {
		words[index] = caser.String(word)
	}
	return strings.Join(words, ".")
}

// Capitalize string following English title capitalization rules.
func Titleize(text string) string {
	words := strings.Fields(text) // Split() + Trim() - empty values.
	last_word := len(words) - 1
	skip_next := false

	for index, word := range words {
		// Check if current iteration should be treated as new sentence.
		skip := bool(skip_next)

		// Check if next iteration should be treated as new sentence.
		skip_next = godis.Contains(end_sentence, word[len(word)-1:])

		// Check if end of sentence or punctuated abbreviation.
		if word[len(word)-1:] == "." {
			// Match "A.Z." (or longer) abbreviation patterns.
			// XXX: Using \pL instead of \w for Unicode letters.
			re := regexp.MustCompile(`^(\pL{1}\.){2,}$`)
			if re.MatchString(word) {
				skip_next = false
			}
		}

		// Check if new sentence, first word or last word first.
		if skip || index == 0 || index == last_word {
			words[index] = Title(word)
		} else if godis.Contains(lowercase, strings.ToLower(word)) {
			words[index] = strings.ToLower(word)
		} else {
			words[index] = Title(word)
		}
	}

	return strings.Join(words, " ")
}
