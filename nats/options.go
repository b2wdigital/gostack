package nats

import (
	"github.com/b2wdigital/goignite/v2/core/config"
)

type Options struct {
	Subjects []string
	Queue    string
}

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
