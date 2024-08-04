package rule

import (
	novelScript "core-engine/internal/core/language/novel-script"
	"core-engine/internal/core/language/novel-script/model"
)

type Chapter struct {
	Keyword string
	Name    string
}

func (c *Chapter) IsValid() bool {
	return len(c.Keyword) > 0 && len(c.Name) > 0
}

func (c *Chapter) ParseRule(nodeId int64, ns *model.NovelScript) error {
	ns.Chapter = c.Name
	return nil
}

func (c *Chapter) ParseWords(words []string) (int, error) {
	for i, word := range words {
		word = removeQuotes(word)

		if i == 0 {
			if word == novelScript.LangKeywordChapter {
				c.Keyword = word
			}
		} else {
			break
		}

		words[i] = word
	}

	value, offset, err := parseQuoteValue(words[1:])
	if err != nil {
		return 0, err
	}

	c.Name = value

	return offset, nil
}
