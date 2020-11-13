package repository

import (
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root          = "gs.repository"
	eventRoot     = root + ".event"
	eventProvider = eventRoot + ".provider"
)

func init() {
	config.Add(eventProvider, "mock", "event provider")
}

func EventProviderValue() string {
	return config.String(eventProvider)
}
