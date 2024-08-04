package rule

import (
	novelScript "core-engine/internal/core/language/novel-script"
	"core-engine/internal/core/language/novel-script/failure"
	"core-engine/internal/core/language/novel-script/model"
	"core-engine/internal/utils"
)

type Person struct {
	Keyword string
	Name    string
	Asset   string
}

func (p *Person) IsValid() bool {
	return len(p.Keyword) > 0 && len(p.Name) > 0 &&
		len(p.Asset) > 0 && utils.IsImage(p.Asset)
}

func (p *Person) ParseWords(words []string) (int, error) {
	for i, word := range words {
		switch i {
		case 0:
			if word != novelScript.LangKeywordPerson {
				return 0, failure.ErrNoPerson
			}
			p.Keyword = word
		case 1:
			p.Name = word
		case 2:
			p.Asset = removeQuotes(words[2])
			break
		}

		words[i] = word
	}

	return 2, nil
}

func (p *Person) ParseRule(nodeId int64, ns *model.NovelScript) error {
	_, ok := ns.Persons[p.Name]
	if ok {
		return failure.ErrPersonAlreadyExist
	}

	ns.Persons[p.Name] = model.Person{
		Name:  p.Name,
		Asset: p.Asset,
	}
	return nil
}
