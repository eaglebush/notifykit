package notifykit

// Sender can be the interface for any Go notification requirement
type Sender interface {
	Send(Message) error
}
