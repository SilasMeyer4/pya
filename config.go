package main

import (
	"encoding/json"
	"log"
	"os"
)

type AppConfiguration struct {
	Repositories       []string
	InstalledLibraries []string
	PersistorVersion   string
}

func CreateInitialConfig(file *os.File) {
	init := AppConfiguration{
		Repositories:       []string{""},
		InstalledLibraries: []string{""},
		PersistorVersion:   "1.0.0",
	}
	if err := json.NewEncoder(file).Encode(init); err != nil {
		log.Fatal("Error encoding initial config:", err)
	}
}

func ReadConfig(file *os.File) AppConfiguration {
	var config AppConfiguration
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		log.Fatal("Error decoding config:", err)
	}
	return config
}
