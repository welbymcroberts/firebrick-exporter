package config

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestParsing(t *testing.T) {
	// Read in config_test_file
	buffer, err := ioutil.ReadFile("config_test_file.yaml")
	if err != nil {
		t.Fatal(err)
	}

	// Parse config
	config, err := ConfigParse(bytes.NewReader(buffer))
	if err != nil {
		t.Fatal(err)
	}

	// Check that config object matches expectations from file
	assert.Equal(t, 1, len(config.Devices), "devices")

	// Check that Attributes exist for first Device
	device1 := config.Devices[0]
	assert.Equal(t, "10.0.0.1", device1.Address, "Device 1: Address")
	assert.Equal(t, "prometheus", device1.Username, "Device 1: Username")
	assert.Equal(t, "prometheus", device1.Password, "Device 1: Password")

	// Check Features for device
	assert.Equal(t, true, device1.Features.Power, "Device 1: Power Feature enabled")

}
