package service

import (
	. "github.com/fazzani/cfcli/domain"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type ConfigService interface {
	Read(config_path string) (*Configuration, error)
}

func Default(config_path string) (*Configuration, error) {
	confBytes, err := ioutil.ReadFile(config_path)

	if err != nil {
		log.Error(err)
	}

	config := Configuration{}
	err = yaml.Unmarshal(confBytes, &config)

	if err != nil {
		log.Error(err)
	}

	return &config, err
}
