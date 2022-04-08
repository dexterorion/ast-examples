package specs

import (
	"encoding/json"
	"os"
)

type Spec struct {
	Package string `json:"package"`
	Files   []File `json:"files"`
}

type File struct {
	Name       string      `json:"name"`
	Structures []Structure `json:"structures"`
}

type Structure struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func NewSpecFromFile(filepath string) (*Spec, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var data Spec
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
