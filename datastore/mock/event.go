package mock

import "github.com/b2wdigital/gostack/repository"

// NewEvent returns a initialized mock
func NewEvent() repository.Event {
	return NewMock()
}
