package publisher

import (
	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/b2wdigital/gostack/cloudevents"
)

const (
	root    = cloudevents.ExtRoot + ".publisher"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable publisher middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
