package rule

import (
	novelScript "core-engine/internal/core/language/novel-script"
	"core-engine/internal/core/language/novel-script/failure"
	"core-engine/internal/core/language/novel-script/model"
	"strings"
)

type PersonReplica struct {
	PersonName string
	Replica    string
}

func (pr *PersonReplica) IsValid() bool {
	return len(pr.PersonName) > 0 && len(pr.Replica) > 0
}

func (pr *PersonReplica) ParseWords(words []string) (int, error) {
	var (
		replica strings.Builder
		isOpen  bool
		offset  int
	)

	pr.PersonName = words[0]

	for _, word := range words[1:] {
		offset++
		if strings.Contains(word, novelScript.LangKeywordQuote) {
			word = strings.ReplaceAll(word, novelScript.LangKeywordQuote, "")

			if !isOpen {
				isOpen = true
				replica.WriteString(word)
				continue
			} else {
				isOpen = false
				replica.WriteString(" " + word)
				break
			}
		}

		replica.WriteString(" " + word)
	}

	pr.Replica = replica.String()

	return offset, nil
}

func (pr *PersonReplica) ParseRule(nodeId int64, ns *model.NovelScript) error {
	_, ok := ns.Persons[pr.PersonName]
	if !ok {
		return failure.ErrPersonNotFound
	}

	ns.Actions[nodeId] = model.Action{
		Replica: &model.Replica{
			PersonName: pr.PersonName,
			Replica:    pr.Replica,
		},
	}

	return nil
}
