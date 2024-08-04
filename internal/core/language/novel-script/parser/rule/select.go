package rule

import (
	novelScript "core-engine/internal/core/language/novel-script"
	"core-engine/internal/core/language/novel-script/failure"
	"core-engine/internal/core/language/novel-script/model"
	"strings"
)

type Select struct {
	Keyword  string
	Variable string
	Variants []string
}

func (s *Select) IsValid() bool {
	return len(s.Keyword) > 0 && len(s.Variants) > 0
}

func (s *Select) ParseWords(words []string) (int, error) {
	var (
		offset       int
		selectStruct strings.Builder
	)

	for i, word := range words {
		switch i {
		case 0:
			if word != novelScript.LangKeywordSelect {
				return 0, failure.ErrSelectSyntax
			}
			offset++
			s.Keyword = word
			continue
		case 1:
			offset++
			s.Variable = word
			continue
		case 2:
			offset++
			if word != novelScript.LangKeywordCondition {
				return 0, failure.ErrSelectSyntax
			}
			continue
		}
		selectStruct.WriteString(word + " ")
	}

	vars := strings.Split(selectStruct.String(), "\"")
	for i, v := range vars {
		v = strings.TrimSpace(v)

		if len(v) == 0 {
			continue
		}
		if i+1 > len(vars)-1 {
			break
		}

		if strings.TrimSpace(vars[i+1]) == novelScript.LangKeywordOrSymbol {
			s.Variants = append(s.Variants, v)
		} else {
			if v == novelScript.LangKeywordOrSymbol {
				continue
			}
			s.Variants = append(s.Variants, v)
			break
		}
	}

	return offset, nil
}

func (s *Select) ParseRule(nodeId int64, ns *model.NovelScript) error {
	ns.Actions[nodeId] = model.Action{
		Select: &model.Select{
			Variable: s.Variable,
			Values:   s.Variants,
		},
	}
	return nil
}
