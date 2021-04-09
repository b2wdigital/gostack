package sns

import (
	"github.com/b2wdigital/goignite/v2/contrib/aws/aws-sdk-go.v2/client/sns"
	"github.com/b2wdigital/gostack/repository"
)

// NewEvent returns a initialized client
func NewEvent(c sns.Client) repository.Event {
	return NewClient(c)
}
