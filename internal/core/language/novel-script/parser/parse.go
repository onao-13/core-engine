package parser

import (
	"bufio"
	novelScript "core-engine/internal/core/language/novel-script"
	"core-engine/internal/core/language/novel-script/failure"
	"core-engine/internal/core/language/novel-script/model"
	"core-engine/internal/utils"
	"github.com/rs/zerolog/log"
	"strings"
)

func (p *Parser) Parse() (*model.NovelScript, error) {
	log.Info().Str("NovelScriptFile", p.file.Name()).Msg("Parse Novel Script file")

	var nodeId int64
	sc := bufio.NewScanner(p.file)
	for sc.Scan() {
		words := strings.Split(sc.Text(), " ")
		if len(words) == 0 {
			continue
		}

		keyword := strings.ToLower(words[0])
		words = words[1:]

		if keyword == novelScript.LangKeywordComment || len(words) == 0 {
			continue
		}

		nodeId++

		switch keyword {
		case novelScript.LangKeywordChapter:
			if err := p.parseChapter(words); err != nil {
				return nil, err
			}
		case novelScript.LangKeywordAct:
			if err := p.parseAct(words); err != nil {
				return nil, err
			}
		case novelScript.LangKeywordPerson:
			if err := p.parsePerson(words); err != nil {
				return nil, err
			}
		case novelScript.LangKeywordBackground:
			if err := p.parseChangeEnvironment(nodeId, words); err != nil {
				return nil, err
			}
		case novelScript.LangKeywordMusic:
			if err := p.parseMusic(nodeId, words); err != nil {
				return nil, err
			}
		case novelScript.LangKeywordUse:
			if err := p.parseUse(words); err != nil {
				return nil, err
			}
		case novelScript.LangKeywordSet:
			if err := p.parseSet(nodeId, words); err != nil {
				return nil, err
			}
		case novelScript.LangKeywordSelect:
			if err := p.parseSelect(nodeId, words); err != nil {
				return nil, err
			}
		case novelScript.LangKeywordIf:

		}

		per, ok := p.out.Persons[keyword]
		if ok {
			if err := p.parsePersonReplica(nodeId, per.Name, words); err != nil {
				return nil, err
			}
			continue
		}
	}

	return p.out, nil
}

func (p *Parser) parseChapter(words []string) error {
	var chapterName string
	chapterName = words[0]

	if len(chapterName) == 0 {
		return failure.ErrNoChapterName
	}

	p.out.Chapter = chapterName

	return nil
}

func (p *Parser) parseAct(words []string) error {
	var actName string
	actName = words[0]
	if len(actName) == 0 {
		return failure.ErrNoActName
	}

	p.out.Act = actName

	return nil
}

func (p *Parser) parsePerson(words []string) error {
	var (
		personName  string
		personAsset string
	)

	if len(words) < 2 {
		return failure.ErrNoPerson
	}

	personName = words[0]
	personAsset = strings.ReplaceAll(words[1], "\"", "")

	if !utils.IsImage(personAsset) {
		return failure.ErrUnsupportedAssetFormat
	}

	if len(personName) == 0 {
		return failure.ErrNoPersonName
	}

	if len(personAsset) == 0 {
		return failure.ErrNoPersonAsset
	}

	p.out.Persons[personName] = model.Person{
		Name:  personName,
		Asset: personAsset,
	}

	return nil
}

func (p *Parser) parsePersonReplica(nodeId int64, perName string, words []string) error {
	var personReplica string
	for _, word := range words {
		personReplica += " " + word
	}

	personReplica = strings.ReplaceAll(personReplica, "\"", "")

	p.out.Actions[nodeId] = model.Action{
		Replica: &model.Replica{
			Replica:    personReplica,
			PersonName: perName,
		},
	}

	return nil
}

func (p *Parser) parseChangeEnvironment(nodeId int64, words []string) error {
	var backgroundAsset string
	backgroundAsset = words[0]

	if len(backgroundAsset) == 0 {
		return failure.ErrParseBackgroundName
	}

	if !utils.IsImage(backgroundAsset) {
		return failure.ErrUnsupportedAssetFormat
	}

	p.out.Actions[nodeId] = model.Action{
		ChangeEnvironment: &model.ChangeEnvironment{
			BackgroundAsset: strings.ReplaceAll(backgroundAsset, "\"", ""),
		},
	}

	return nil
}

func (p *Parser) parseMusic(nodeId int64, words []string) error {
	var musicName string
	musicName = words[0]
	if len(musicName) == 0 {
		return failure.ErrParseMusicName
	}

	p.out.Actions[nodeId] = model.Action{
		ChangeEnvironment: &model.ChangeEnvironment{
			MusicFile: strings.ReplaceAll(musicName, "\"", ""),
		},
	}

	return nil
}

func (p *Parser) parseUse(words []string) error {
	var useType string
	useType = words[0]

	if len(useType) == 0 {
		return failure.ErrParseUseType
	}
	words = words[1:]

	switch useType {
	case novelScript.TypeEns:
		return p.parseUseEns(words)
	}

	return failure.ErrUseTypeNotFound
}

func (p *Parser) parseUseEns(words []string) error {
	var (
		ensName string
		ensFile string
	)

	if len(words) < 2 {
		return failure.ErrParseEnsSyntax
	}

	ensName = words[0]
	ensFile = strings.ReplaceAll(words[1], "\"", "")
	if len(ensName) == 0 {
		return failure.ErrNoEnsName
	}

	if len(ensFile) == 0 {
		return failure.ErrNoEnsFile
	}

	p.out.EnsFiles[ensName] = ensFile

	return nil
}

func (p *Parser) parseSet(nodeId int64, words []string) error {
	var setType string
	setType = words[0]
	if len(setType) == 0 {
		return failure.ErrNoSetName
	}

	words = words[1:]
	switch setType {
	case novelScript.TypeEns:
		return p.parseSetEns(nodeId, words)
	}

	return failure.ErrSetTypeNotFound
}

func (p *Parser) parseSetEns(nodeId int64, words []string) error {
	var (
		ensName  string
		ensKey   string
		ensValue string
	)

	if len(words) < 2 {
		return failure.ErrParseUseEnsSyntax
	}

	splitEnsNameKey := strings.Split(words[0], novelScript.LangKeywordDot)
	if len(splitEnsNameKey) != 2 {
		return failure.ErrParseUseEnsSyntax
	}

	ensName = splitEnsNameKey[0]
	ensKey = splitEnsNameKey[1]

	if words[1] != novelScript.LangKeywordCondition {
		return failure.ErrParseUseEnsSyntax
	}

	ensValue = words[2]

	p.out.Actions[nodeId] = model.Action{
		ChangeEnsValue: &model.ChangeEnsValue{
			Name:  ensName,
			Key:   ensKey,
			Value: ensValue,
		},
	}

	return nil
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
		selectValue  string
	)

	for _, word := range words[2:] {
		switch {
		case word == novelScript.LangKeywordOrSymbol:
			selectValue = strings.Replace(selectValue, " ", "", 1)
			selectValues = append(selectValues, strings.ReplaceAll(selectValue, "\"", ""))
			selectValue = ""
			continue
		case strings.Contains(word, novelScript.LangKeywordQuote):
			selectValue += " " + word
			continue
		}
	}
	selectValue = strings.Replace(selectValue, " ", "", 1)
	selectValues = append(selectValues, strings.ReplaceAll(selectValue, "\"", ""))
	selectValue = " "

	p.out.Actions[nodeId] = model.Action{
		Select: &model.Select{
			Variable: selectVariable,
			Values:   selectValues,
		},
	}

	return nil
}

//func (p *Parser) parseCondition(nodeId int64, words []string) error {
//
//}
