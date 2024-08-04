package project

import (
	"core-engine/internal/common"
	"core-engine/internal/utils"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
)

var (
	TemplateProjectFiles = []string{
		"README.md", "example.ns", "about.json",
	}
	TemplateProjectFolders = []string{
		common.FolderScenarios,
	}
)

func (p Project) CreateTemplateFiles() error {
	log.Info().Msg("Create Template Files")

	for _, file := range TemplateProjectFiles {
		fName := filepath.Join(p.Location, file)
		if utils.FileIsExist(fName) {
			continue
		}

		f, err := os.Create(fName)
		if err != nil {
			log.Error().Err(err).Msg("Error creating template file")
			return err
		}

		log.Info().Str("file", f.Name()).Msg("Created template file")
	}

	for _, folder := range TemplateProjectFolders {
		fName := filepath.Join(p.Location, folder)
		if utils.FileIsExist(fName) {

		}
		if err := os.Mkdir(fName, os.ModePerm); err != nil {
			log.Error().Err(err).Msg("Error creating template folder")
			return err
		}
	}

	log.Info().Msg("Success create template files")
	return nil
}
