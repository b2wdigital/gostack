package cloudevents

import (
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root                  = "gs.cloudevents"
	ExtRoot               = root + ".ext"
	handleDiscardEventsID = root + ".handle.discard.ids"
)

func init() {
	config.Add(handleDiscardEventsID, "", "cloudevents events id that will not be processed, comma separated")
}

func HandleDiscardEventsIDValue() string {
	return config.String(handleDiscardEventsID)
}
