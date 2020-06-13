package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Email   string `yaml:"email"`
	Zoomurl string `yaml:"zoomurl"`
}

var config Config

func readconfig(path string) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(b, &config)
	if err != nil {
		return err
	}

	return nil
}
