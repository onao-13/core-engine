package project

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
)

type Project struct {
	Name        string
	Description string
	Location    string
	IsEmpty     bool
}

func NewProject(name string, desc string, location string, isEmpty bool) *Project {
	log.Info().Msgf("Create new project: %s into %s", name, location)

	p := &Project{
		Name: name,
		Description: lo.
			If(len(desc) != 0, desc).
			Else(fmt.Sprintf("This is amazing %s project!", name)),
		Location: location,
		IsEmpty:  isEmpty,
	}

	if err := p.IsValid(); err != nil {
		log.Error().Err(err).Msg("Invalid project field")
		return nil
	}

	return p
}

func (p Project) IsValid() error {
	switch {
	case len(p.Name) == 0:
		return fmt.Errorf("project name is empty")
	case len(p.Location) == 0:
		return fmt.Errorf("project location is empty")
	}
	return nil
}
