package config

import (
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
	// ensure no prexisting config
	configDir, err := os.UserConfigDir()
	if err != nil {
		t.Errorf("ERROR: %s\n", err.Error())
		t.FailNow()
	}
	os.Remove(configDir + "/ocelot.toml")

	var config Config

	err = config.Read()

	if err != nil {
		t.Errorf("ERROR: %s\n", err.Error())
		t.FailNow()
	}
	if config.SecretKey == "" {
		t.Errorf("ERROR: Secret key not generated correctly.\n")
		t.FailNow()
	}
}

func TestPrexistingConfig(t *testing.T) {
	var config Config

	err := config.Read()

	if err != nil {
		t.Errorf("ERROR: %s\n", err.Error())
		t.FailNow()
	}
	if config.SecretKey == "" {
		t.Errorf("ERROR: Secret key not generated correctly.\n")
		t.FailNow()
	}
}
