package scenario

import (
	"core-engine/internal/common"
	"core-engine/internal/core/language/novel-script/model"
	"core-engine/internal/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"time"
)

type Meta struct {
	CreatedAt time.Time `json:"createdAt"`
	Act       string    `json:"act"`
	Chapter   string    `json:"chapter"`
}

type Scenario struct {
	Meta  Meta   `json:"meta"`
	Nodes []Node `json:"nodes"`
}

func NewScenario(ns *model.NovelScript) *Scenario {
	if ns == nil {
		return nil
	}

	return ParseNovelScriptFile(ns)
}

func (s *Scenario) Save(projectPath string) error {
	log.Info().Msg("Save scenario file")

	fPath := filepath.Join(projectPath, common.FolderScenarios, fmt.Sprintf("%s.json", s.Name()))

	var (
		f   *os.File
		err error
	)

	if utils.FileIsExist(fPath) {
		f, err = os.OpenFile(fPath, os.O_WRONLY, os.ModePerm)
	} else {
		f, err = os.Create(fPath)
		if err != nil {
			log.Error().Err(err).Msg("Failed to create file")
			return err
		}
	}

	defer f.Close()

	scen, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal json")
		return err
	}

	if err = os.Truncate(fPath, 0); err != nil {
		log.Error().Err(err).Msg("Failed to truncate file")
		return err
	}

	_, err = f.Write(scen)
	if err != nil {
		log.Error().Err(err).Msg("Failed to write json")
		return err
	}

	log.Info().Msg("Save scenario file successfully")

	return nil
}

func (s *Scenario) Name() string {
	hash := base64.StdEncoding.EncodeToString([]byte(s.Meta.Act + s.Meta.Chapter))
	return fmt.Sprintf("scenario_%s_%s_%s", s.Meta.Act, s.Meta.Chapter, hash)
}
