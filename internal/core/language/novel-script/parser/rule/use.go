package rule

import (
	novelScript "core-engine/internal/core/language/novel-script"
	"core-engine/internal/core/language/novel-script/failure"
	"core-engine/internal/core/language/novel-script/model"
)

type UseEns struct {
	Name string
	File string
}

type Use struct {
	Keyword string
	Type    string
	Ens     *UseEns
}

func (u *Use) IsValid() bool {
	return len(u.Keyword) > 0 && len(u.Type) > 0
}

func (u *Use) ParseWords(words []string) (int, error) {
	for i, word := range words {

		switch i {
		case 0:
			if word != novelScript.LangKeywordUse {
				return 0, failure.ErrSyntaxUse
			}
			u.Keyword = word
		case 1:
			if word != novelScript.TypeEns {
				return 0, failure.ErrSyntaxUse
			}
			u.Type = word
		case 2:
			break
		}

		words[i] = word
	}

	switch u.Type {
	case novelScript.TypeEns:
		u.Ens = &UseEns{}

		u.Ens.Name = words[2]
		u.Ens.File = removeQuotes(words[3])
	}

	return 3, nil
}

func (u *Use) ParseRule(nodeId int64, ns *model.NovelScript) error {
	switch u.Type {
	case novelScript.TypeEns:
		_, ok := ns.EnsFiles[u.Ens.Name]
		if ok {
			return failure.ErrEnsAlreadyExist
		}

		ns.EnsFiles[u.Ens.Name] = u.Ens.File
	}

	return nil
}
