package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Root struct {
	Base  string
	Parts []PartGroup
}

type PartGroup struct {
	Group   string
	Details []struct {
		Name string
		Text string
	}
}

func loadYaml(path string) (*Root, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	r := Root{}
	if err := yaml.Unmarshal(b, &r); err != nil {
		return nil, err
	}

	return &r, nil
}
