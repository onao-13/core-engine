package rule

import (
	novelScript "core-engine/internal/core/language/novel-script"
	"core-engine/internal/core/language/novel-script/failure"
	"core-engine/internal/core/language/novel-script/model"
	"core-engine/internal/utils"
)

type Background struct {
	Keyword string
	Asset   string
}

func (b *Background) IsValid() bool {
	return len(b.Keyword) > 0 && len(b.Asset) > 0 &&
		utils.IsImage(b.Asset)
}

func (b *Background) ParseWords(words []string) (int, error) {
	for i, word := range words {

		switch i {
		case 0:
			if word != novelScript.LangKeywordBackground {
				return 0, failure.ErrSyntaxBackground
			}
			b.Keyword = word
		case 1:
			b.Asset = removeQuotes(word)
			break

		}

		words[i] = word
	}

	return 1, nil
}

func (b *Background) ParseRule(nodeId int64, ns *model.NovelScript) error {
	ns.Actions[nodeId] = model.Action{
		ChangeEnvironment: &model.ChangeEnvironment{
			BackgroundAsset: b.Asset,
		},
	}
	return nil
}
