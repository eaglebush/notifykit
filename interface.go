package notifykit

// Message can be the interface for any Go message requirement
type Message interface {
	// Encodes the plain text message body to base64 for safe JSON value
	SetMsg(string) error

	// Decodes the encoded message body from base64
	GetMsg() string
}

// Sender can be the interface for any Go notification requirement
type Sender interface {
	Send(Message) error
}
