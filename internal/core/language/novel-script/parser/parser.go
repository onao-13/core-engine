package parser

import (
	"core-engine/internal/core/language/novel-script/model"
	"github.com/rs/zerolog/log"
	"os"
)

type Parser struct {
	file *os.File
	out  *model.NovelScript
}

func (p *Parser) Load(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 077)
	if err != nil {
		log.Fatal().Err(err).Msg("Error opening file")
		return err
	}

	if p.file != nil {
		if err := p.file.Close(); err != nil {
			log.Error().Err(err).Msg("failed to close file")
		}
	}

	p.file = file
	p.out = &model.NovelScript{
		Persons:  make(map[string]model.Person),
		Actions:  make(map[int64]model.Action),
		EnsFiles: make(map[string]string),
	}

	return nil
}
