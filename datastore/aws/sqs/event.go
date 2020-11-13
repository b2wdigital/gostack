package sqs

import (
	giawsclientsqs "github.com/b2wdigital/goignite/v2/contrib/aws/aws-sdk-go.v2/client/sqs"
	"github.com/b2wdigital/gostack/repository"
)

// NewEvent returns a initialized client
func NewEvent(c giawsclientsqs.Client) repository.Event {
	return NewClient(c)
}
