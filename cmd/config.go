package cmd

import (
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root = "gs.cmd"
	def  = root + ".default"
)

func init() {
	config.Add(def, "", "default cmd")
}

func Default() string {
	return config.String(def)
}
