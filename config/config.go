package config

import (
	"log"
	"os"
	"path"

	"github.com/yusuftatli/hepsiburada/models"
	"gopkg.in/yaml.v2"
)

func NewDefaultConfig() (*models.ProviderConfig, error) {
	configFilePath, err := getFilePath("config.yml")
	if err != nil {
		return nil, err
	}

	file, err := os.Open(configFilePath)
	if err != nil {
		log.Panicln("Failed to open config.yml file", err)
		// errors.Wrap(err, "Failed to open config.yml file")
		return nil, err
	}

	defer file.Close()

	cfg := new(models.ProviderConfig)
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Panicln("Failed to decode yaml config", err)
	}

	return cfg, nil
}

func getFilePath(fileName string) (string, error) {
	rootPath, err := os.Getwd()
	if err != nil {
		// errors.Wrap(err, "Failed to get root path")
		return "", err
	}

	return path.Join(path.Dir(rootPath), "/hepsiburada/"+fileName), nil
}
