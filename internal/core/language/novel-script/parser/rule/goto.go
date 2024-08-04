package rule

import (
	novelScript "core-engine/internal/core/language/novel-script"
	"core-engine/internal/core/language/novel-script/failure"
	"core-engine/internal/core/language/novel-script/model"
)

type Goto struct {
	Keyword string
	File    string
}

func (g *Goto) IsValid() bool {
	return len(g.Keyword) > 0 && len(g.File) > 0
}

func (g *Goto) ParseWords(words []string) (int, error) {
	for i, word := range words {
		switch i {
		case 0:
			if word != novelScript.LangKeywordGoto {
				return 0, failure.ErrSyntaxGoto
			}
			g.Keyword = word
		case 1:
			g.File = removeQuotes(word)
			break
		}
	}

	return 1, nil
}

func (g *Goto) ParseRule(nodeId int64, ns *model.NovelScript) error {
	ns.Actions[nodeId] = model.Action{
		Goto: &model.Goto{
			File: g.File,
		},
	}
	return nil
}
