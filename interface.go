package notifykit

// Message can be the interface for any Go message requirement
type Message interface {
	// Encodes the plain text message to base64 for safe JSON value
	SetMsg(string) error
}

// Sender can be the interface for any Go notification requirement
type Sender interface {
	Send(Message) error
}
