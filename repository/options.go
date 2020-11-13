package repository

import "github.com/b2wdigital/goignite/v2/core/config"

type Options struct {
	Provider string
}

func DefaultEventOptions() (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath(eventRoot, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
