package notifykit

import (
	"errors"
	"slices"
)

// Sender can be the interface for any Go notification requirement
type Sender interface {
	Send(Message) error
	Messages() []string
}

var implSender func(...any) Sender

// Register the implementation
func Register(impl func(...any) Sender) {
	implSender = impl
}

// New retrieves the Sender driver
func New(host string, opts ...Option) (Sender, error) {
	if implSender == nil {
		return nil, errors.New("sender: no sender registered")
	}

	opts = slices.Insert(opts, 0, Host(host))

	// Unpack the options into an slice of raw data
	var driverArgs []any
	for _, opt := range opts {
		opt(&driverArgs)
	}

	// Pass the slice directly down as variadic arguments to the driver
	return implSender(driverArgs...), nil
}
