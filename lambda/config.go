package lambda

import (
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root = "gs.lambda"
	skip = root + ".skip"
)

func init() {
	config.Add(skip, false, "skip all triggers")
}

func SkipValue() bool {
	return config.Bool(skip)
}
