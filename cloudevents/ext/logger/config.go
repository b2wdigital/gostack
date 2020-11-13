package logger

import (
	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/b2wdigital/gostack/cloudevents"
)

const (
	root    = cloudevents.ExtRoot + ".logger"
	enabled = root + ".enabled"
	level   = root + ".level"
)

func init() {
	config.Add(enabled, true, "enable/disable logger middleware")
	config.Add(level, "INFO", "sets log level INFO/DEBUG/TRACE")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func Level() string {
	return config.String(level)
}
