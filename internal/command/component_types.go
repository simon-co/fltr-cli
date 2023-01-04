package command

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/simon-co/fltr-cli/internal/apperr"
)

type ComponentType string

type ComponentTypes struct {
	View,
	Dialog,
  Navigator ComponentType
}

type ComponentTypeMap map[string]any

func (self ComponentTypes) ToMap() ComponentTypeMap {
	typeMap := ComponentTypeMap{}
	jsonStr, _ := json.Marshal(self)
	json.Unmarshal(jsonStr, &typeMap)
	return typeMap
}

func (self ComponentTypeMap) ToList() []string {
	l := []string{}
	for k := range self {
		l = append(l, k)
	}
	return l
}

func (self ComponentType) toBlueprint(output io.Writer) (ComponentBlueprint, error) {
	var blueprint ComponentBlueprint
	switch self {
	case "View":
		blueprint = ViewBlueprint{}
  case "Dialog":
    blueprint = DialogBlueprint{}
	default:
		return nil, errors.New("Invalid component type")
	}
	blueprint, err := blueprint.New(output)
  if err != nil {
		return nil, apperr.Parse(err)
	}
	return blueprint, nil
}
