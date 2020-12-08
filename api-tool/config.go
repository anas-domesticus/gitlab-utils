package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

type config struct {
	Token string
}

func NewConfig() *config {
	envVars := [...]string{"GITLABTOKEN"}
	var fatalError = false
	for _, v := range envVars {
		_, varPresent := os.LookupEnv(v)
		if !(varPresent) {
			log.Error("Missing " + v + " environment variable")
			fatalError = true
		}
	}
	if fatalError {
		os.Exit(1)
	}

	c := config{
		Token: os.Getenv("GITLABTOKEN"),
	}

	return &c
}
