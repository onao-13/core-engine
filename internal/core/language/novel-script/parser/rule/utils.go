package rule

import (
	novelScript "core-engine/internal/core/language/novel-script"
	"github.com/samber/lo"
	"strings"
)

func parseQuoteValue(words []string) (value string, offset int, err error) {
	var isOpen = false

	for _, word := range words {
		if strings.Contains(word, novelScript.LangKeywordQuote) {
			if !isOpen {
				isOpen = true
				offset++
				value = strings.ReplaceAll(word, novelScript.LangKeywordQuote, "")
				continue
			} else {
				offset++
				value += " " + strings.ReplaceAll(word, novelScript.LangKeywordQuote, "")
				isOpen = false
				break
			}
		}

		if isOpen {
			offset++
			value += " " + word
		}
	}

	return
}

func removeQuotes(word string) string {
	return lo.
		If(
			strings.Contains(word, novelScript.LangKeywordQuote),
			strings.ReplaceAll(word, novelScript.LangKeywordQuote, ""),
		).
		Else(word)
}
