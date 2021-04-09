package util

import (
	"encoding/json"

	"github.com/b2wdigital/goignite/v2/core/errors"
	v2 "github.com/cloudevents/sdk-go/v2"
)

func JSONBytes(event v2.Event) ([]byte, error) {

	rawMessage, err := json.Marshal(event)
	if err != nil {
		return nil, errors.Wrap(err, errors.Errorf("error on json marshal. %s", err.Error()))
	}

	return rawMessage, nil
}
