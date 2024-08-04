package parser

import (
	"bufio"
	novelScript "core-engine/internal/core/language/novel-script"
	"core-engine/internal/core/language/novel-script/failure"
	"core-engine/internal/core/language/novel-script/model"
	"core-engine/internal/core/language/novel-script/parser/rule"
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
)

var KeyParser = map[string]rule.Parse{
	novelScript.LangKeywordChapter:    &rule.Chapter{},
	novelScript.LangKeywordAct:        &rule.Act{},
	novelScript.LangKeywordPerson:     &rule.Person{},
	novelScript.LangKeywordBackground: &rule.Background{},
	novelScript.LangKeywordMusic:      &rule.Music{},
	novelScript.LangKeywordSet:        &rule.Set{},
	novelScript.LangKeywordUse:        &rule.Use{},
	novelScript.LangKeywordSelect:     &rule.Select{},
	novelScript.LangKeywordGoto:       &rule.Goto{},
}

func (p *Parser) Parse() (*model.NovelScript, error) {
	log.Info().Str("NovelScriptFile", p.file.Name()).Msg("Parse Novel Script file")

	sc := bufio.NewScanner(p.file)
	var (
		file   []string
		nodeId int64
		i      int
		offset int
		err    error
	)

	for sc.Scan() {
		line := sc.Text()
		if !strings.Contains(line, novelScript.LangKeywordComment) {
			file = append(file, strings.Split(line, " ")...)
		}
	}

	for _, word := range file {
		word = file[i]
		word = strings.TrimSpace(word)
		word = strings.ToLower(word)
		if i < len(file)-2 {
			file[i] = word
		} else {
			break
		}

		parser, ok := KeyParser[word]
		if !ok {
			if p.isPerson(word) {
				parser = &rule.PersonReplica{}
			} else {
				i++
				continue
			}
		}

		nodeId++
		offset, err = parser.ParseWords(file[i:])
		if err != nil {
			return nil, err
		}

		if !parser.IsValid() {
			return nil, fmt.Errorf("is not a valid novel script")
		}

		if err = parser.ParseRule(nodeId, p.out); err != nil {
			return nil, err
		}

		i += offset + 1
	}

	return p.out, nil
}

func (p *Parser) isPerson(word string) bool {
	_, ok := p.out.Persons[word]
	return ok
}

func (p *Parser) parseSelect(nodeId int64, words []string) error {
	var selectVariable string
	selectVariable = words[0]
	if len(selectVariable) == 0 {
		return failure.ErrSelectSyntaxError
	}

	if words[1] != novelScript.LangKeywordCondition {
		return failure.ErrSelectSyntaxError
	}

	var (
		selectValues = make([]string, 0)
		line         string
	)

	words = words[2:]
	for _, word := range words {
		if strings.Contains(word, novelScript.LangKeywordQuote) {
			word = strings.ReplaceAll(word, novelScript.LangKeywordQuote, "")
		}
		line += word + " "
	}

	selectValues = strings.Split(line, fmt.Sprintf(" %s ", novelScript.LangKeywordOrSymbol))

	p.out.Actions[nodeId] = model.Action{
		Select: &model.Select{
			Variable: selectVariable,
			Values:   selectValues,
		},
	}

	return nil
}
