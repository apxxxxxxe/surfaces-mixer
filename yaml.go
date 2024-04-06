package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Root struct {
	Base  string      `yaml:"base"`
	Parts []GroupData `yaml:"parts"`
	Raw   string      `yaml:"raw"`
}

type GroupData struct {
	Name  string     `yaml:"group"`
	Poses []PoseData `yaml:"details"`
}

type PoseData struct {
	Name string `yaml:"name"`
	Text string `yaml:"text"`
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
