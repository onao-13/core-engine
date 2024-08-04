package rule

import (
	"core-engine/internal/common"
	novelScript "core-engine/internal/core/language/novel-script"
	"core-engine/internal/core/language/novel-script/failure"
	"core-engine/internal/core/language/novel-script/model"
	"core-engine/internal/utils"
)

type Music struct {
	Keyword string
	Asset   string
}

func (m *Music) IsValid() bool {
	return len(m.Keyword) > 0 && len(m.Asset) > 0 &&
		utils.IsMedia(m.Asset, common.MediaFormatMp3)
}

func (m *Music) ParseWords(words []string) (int, error) {
	for i, word := range words {
		switch i {
		case 0:
			if word != novelScript.LangKeywordMusic {
				return 0, failure.ErrSyntaxMusic
			}
			m.Keyword = word
		case 1:
			m.Asset = removeQuotes(word)
			break
		}

		words[i] = word
	}

	//asset, offset, err := parseQuoteValue(words[1:])
	//if err != nil {
	//	return 0, err
	//
	//}
	//
	//m.Asset = asset

	return 2, nil
}

func (m *Music) ParseRule(nodeId int64, ns *model.NovelScript) error {
	ns.Actions[nodeId] = model.Action{
		ChangeEnvironment: &model.ChangeEnvironment{
			MusicFile: m.Asset,
		},
	}
	return nil
}
