package nats

import (
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root     = "gs.nats"
	subjects = root + ".subjects"
	queue    = root + ".queue"
)

func init() {
	config.Add(subjects, []string{"changeme"}, "nats listener subjects")
	config.Add(queue, "changeme", "nats listener queue")
}

func SubjectsValue() []string {
	return config.Strings(subjects)
}

func QueueValue() string {
	return config.String(queue)
}
