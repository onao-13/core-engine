package rule

import (
	novelScript "core-engine/internal/core/language/novel-script"
	"core-engine/internal/core/language/novel-script/model"
)

type Act struct {
	Keyword string
	Name    string
}

func (a *Act) IsValid() bool {
	return len(a.Keyword) > 0 && len(a.Name) > 0
}

func (a *Act) ParseWords(words []string) (int, error) {
	for i, word := range words {
		word = removeQuotes(word)

		if i == 0 {
			if word == novelScript.LangKeywordAct {
				a.Keyword = word
			}
		} else {
			break
		}

		words[i] = word
	}

	name, offset, err := parseQuoteValue(words)
	if err != nil {
		return 0, err
	}

	a.Name = name

	return offset, nil
}

func (a *Act) ParseRule(nodeId int64, ns *model.NovelScript) error {
	ns.Act = a.Name
	return nil
}
