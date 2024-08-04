package rule

import (
	novelScript "core-engine/internal/core/language/novel-script"
	"core-engine/internal/core/language/novel-script/failure"
	"core-engine/internal/core/language/novel-script/model"
	"strings"
)

type SetEns struct {
	Name     string
	Variable string
	Value    string
}

type Set struct {
	Keyword string
	Type    string
	Ens     *SetEns
}

func (s *Set) IsValid() bool {
	return len(s.Keyword) > 0 && len(s.Type) > 0
}

func (s *Set) ParseWords(words []string) (int, error) {
	var offset int

	for i, word := range words {

		switch i {
		case 0:
			if word != novelScript.LangKeywordSet {
				return 0, failure.ErrSyntaxSet
			}
			s.Keyword = word
		case 1:
			if word != novelScript.TypeEns {
				return 0, failure.ErrSyntaxSet
			}
			s.Type = word
		case 2:
			words = words[2:]

			switch s.Type {
			case novelScript.TypeEns:
				s.Ens = &SetEns{}

				splitEnsNameKey := strings.Split(removeQuotes(words[0]), novelScript.LangKeywordDot)
				if len(splitEnsNameKey) != 2 {
					return 0, failure.ErrParseUseEnsSyntax
				}

				s.Ens.Name = splitEnsNameKey[0]
				s.Ens.Variable = splitEnsNameKey[1]

				if words[1] != novelScript.LangKeywordCondition {
					return 0, failure.ErrParseUseEnsSyntax
				}

				s.Ens.Value = removeQuotes(words[2])

				offset = 4
			}
			return offset, nil
		}

		words[i] = word
	}

	return offset, nil
}

func (s *Set) ParseRule(nodeId int64, ns *model.NovelScript) error {
	switch s.Type {
	case novelScript.TypeEns:
		_, ok := ns.EnsFiles[s.Ens.Name]
		if !ok {
			return failure.ErrNoEnsFile
		}

		ns.Actions[nodeId] = model.Action{
			ChangeEnsValue: &model.ChangeEnsValue{
				Name:  s.Ens.Name,
				Key:   s.Ens.Variable,
				Value: s.Ens.Value,
			},
		}
	}
	return nil
}
