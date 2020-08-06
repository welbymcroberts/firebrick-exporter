package config

import (
	"bytes"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
)

//// Features that are enabled or disabled on a per device basis
type FeatureConfig struct {
	Power    bool `yaml:"power,omitempty"`
	Threads  bool `yaml:"threads,omitempty"`
	Subnets  bool `yaml:"subnets,omitempty"`
	DNS      bool `yaml:"dns,omitempty"`
	DHCP     bool `yaml:"dhcp,omitempty"`
	Profiles bool `yaml:"profiles,omitempty"`
	PPPoE    bool `yaml:"pppoe,omitempty"`
	L2TP     bool `yaml:"l2tp,omitempty"`
	Sessions bool `yaml:"sessions,omitempty"`
}

//// Individual Devices
type DeviceConfig struct {
	Address  string        `yaml:"address"`
	Username string        `yaml:"username,omitempty"`
	Password string        `yaml:"password,omitempty"`
	Features FeatureConfig `yaml:"features,omitempty"`
}

//// Exporter Wide
type Config struct {
	Devices []*DeviceConfig `yaml:"devices,omitempty"`
}

func ConfigParse(r io.Reader) (*Config, error) {
	// read everything from io.Reader
	buffer, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	// Create config instance from yaml
	config := &Config{}
	err = yaml.Unmarshal(buffer, config)
	if err != nil {
		return nil, err
	}

	// TODO: Set default features if not set on device
	return config, nil
}

func ConfigLoadFromFile() (*Config, error) {
	// Load from file
	buffer, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	// Parse config
	config, err := ConfigParse(bytes.NewReader(buffer))
	if err != nil {
		return nil, err
	}

	return config, nil
}
