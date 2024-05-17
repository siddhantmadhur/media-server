package config

import (
	"os"
	"runtime"
)

type ServerInformation struct {
	Hostname        string `json:"hostname"`
	ServerVersion   string `json:"server_version"`
	OperatingSystem string `json:"operating_system"`
	FinishedWizard  bool   `json:"finished_wizard"`
}

func getServerInformation() (ServerInformation, error) {
	var config Config
	err := config.Read()
	if err != nil {
		return ServerInformation{}, err
	}
	hostname, err := os.Hostname()
	if err != nil {
		return ServerInformation{}, err
	}
	return ServerInformation{
		Hostname:        hostname,
		ServerVersion:   os.Getenv("ocelot_version"),
		OperatingSystem: runtime.GOOS,
		FinishedWizard:  config.FinishedWizard,
	}, nil

}
