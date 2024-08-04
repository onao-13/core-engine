package scenario

import (
	"core-engine/internal/core/language/novel-script/model"
	"github.com/rs/zerolog/log"
	"sort"
	"time"
)

func ParseNovelScriptFile(ns *model.NovelScript) *Scenario {
	log.Info().Msg("Parse novel script file to scenario file")

	scen := &Scenario{
		Meta: Meta{
			CreatedAt: time.Now(),
			Act:       ns.Act,
			Chapter:   ns.Chapter,
		},
		Nodes: make([]Node, 0),
	}

	ids := make([]int, 0)
	for id := range ns.Actions {
		ids = append(ids, int(id))
	}

	sort.Ints(ids)

	for _, id := range ids {
		action, ok := ns.Actions[int64(id)]
		if !ok {
			continue
		}
		switch {
		case action.ChangeEnvironment != nil:
			scen.Nodes = append(scen.Nodes, Node{
				EnvironmentInfo: &EnvironmentInfo{
					Background: action.ChangeEnvironment.BackgroundAsset,
					Music:      action.ChangeEnvironment.MusicFile,
				},
			})
		case action.Replica != nil:
			per, ok := ns.Persons[action.Replica.PersonName]
			if !ok {
				continue
			}

			scen.Nodes = append(scen.Nodes, Node{
				PersonInfo: &PersonReplicaInfo{
					Person:  per.Name,
					Asset:   per.Asset,
					Replica: action.Replica.Replica,
				},
			})
		case action.ChangeEnsValue != nil:
			ensFile, ok := ns.EnsFiles[action.ChangeEnsValue.Name]
			if !ok {
				continue
			}

			scen.Nodes = append(scen.Nodes, Node{
				Action: &Action{
					ChangeEns: &ActionChangeEns{
						NewValue: action.ChangeEnsValue.Value,
						Key:      action.ChangeEnsValue.Key,
						File:     ensFile,
					},
				},
			})
		}
	}

	log.Info().Msg("Successfully parsed novel script file to scenario file")

	return scen
}
