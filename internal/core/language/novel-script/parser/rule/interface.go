package rule

import "core-engine/internal/core/language/novel-script/model"

type Parse interface {
	IsValid() bool
	ParseWords(words []string) (int, error)
	ParseRule(nodeId int64, ns *model.NovelScript) error
}
