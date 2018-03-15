package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

//
// User ...
//
type User struct {
	User   string `yaml:"user"`
	APIKey string `yaml:"apikey"`
	Secret string `yaml:"secret"`
}

//
// Config ...
//
type Config struct {
	Endpoint string `yaml:"endpoint"`
	Language string `yaml:"language"`
	Login    []User `yaml:"login"`
}

//
// ReadConfigFile reads a configuration file and stores
// it into a Config struct
//
func ReadConfigFile(filepath string) (*Config, error) {
	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		return nil, err
	}

	config := Config{}
	err = yaml.Unmarshal(data, &config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}

//
// SaveConfigFile stores the configuration on disc
//
func SaveConfigFile(filepath string, cfg *Config) error {
	out, err := yaml.Marshal(cfg)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath, out, 0660)

	if err != nil {
		return err
	}

	return nil
}
