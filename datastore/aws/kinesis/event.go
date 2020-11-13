package kinesis

import (
	"github.com/b2wdigital/goignite/v2/contrib/aws/aws-sdk-go.v2/client/kinesis"
	"github.com/b2wdigital/gostack/repository"
)

// NewEvent returns a initialized client
func NewEvent(c kinesis.Client, options *Options) repository.Event {
	return NewClient(c, options)
}
