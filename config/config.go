package config

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
)

func Get(key string) (string, error) {
	return get(key, "")
}

func get(key string, defaultValue string) (string, error) {
	value, isPresent := os.LookupEnv(key)

	if !isPresent {
		log.Warn("Unable to find configuration for key " + key)
		return defaultValue, errors.New("Unable to find configuration for key " + key)
	}

	return value, nil
}
