package project

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"time"
)

func (p Project) Create() error {
	log.Info().Msg("Start creating project process")

	if err := p.CreateTemplateFiles(); err != nil {
		log.Error().Err(err).Msg("Failed to create template files")
		return err
	}

	if err := p.writeAboutInfo(); err != nil {
		log.Error().Err(err).Msg("Failed to create template files")
		return err
	}

	log.Info().Msg("Success finished creating project process")

	return nil
}

func (p Project) writeAboutInfo() error {
	log.Info().Msg("Writing about project info")

	fAbout, err := os.OpenFile(filepath.Join(p.Location, "about.json"), os.O_WRONLY, 0777)
	if err != nil {
		log.Error().Err(err).Msg("Failed to open about.json")
		return err
	}

	defer fAbout.Close()

	about := About{
		Name:        p.Name,
		Version:     "1.0.0",
		Author:      "",
		Email:       "",
		Description: p.Description,
		DateCreated: time.Now(),
	}

	aboutInfo, err := json.MarshalIndent(about, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal about info")
		return err
	}

	if _, err = fAbout.Write(aboutInfo); err != nil {
		log.Error().Err(err).Msg("Failed to write about info")
		return err
	}

	log.Info().Msg("Success writing about info")

	return nil
}
