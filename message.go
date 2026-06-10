package notifykit

import "errors"

// Role
//
// SENDER, FROM and REPLYTO depends on the library implementing this interface
const (
	TO Role = iota
	CC
	BCC
	SENDER  // Depends on the library using the interface
	FROM    // Depends on the library using the interface
	REPLYTO // Depends on the library using the interface
)

// ResourceType
const (
	ATTACHMENT ResourceType = 0
	INLINE     ResourceType = 1
)

var (
	ErrMailHasNoApplicationId          error = errors.New("mail has no application id")
	ErrMailFileIsNotProvided           error = errors.New("file is not provided")
	ErrMailFileDoesNotExist            error = errors.New("file does not exist")
	ErrMailRecipientNameNotProvided    error = errors.New("recipient name is not provided")
	ErrMailRecipientAddressNotProvided error = errors.New("recipient address is not provided")
	ErrMailMessageNotProvided          error = errors.New("message is not provided")
)

type (
	Role         uint8
	ResourceType uint8

	// Message interface definition for Sender compatibility use
	Message interface {
		// AddRecipient adds a new mail recipient with role.
		//
		// The address parameter is the e-mail address. But it can be phone number, or any social media handler
		AddRecipient(name string, address string, role Role) error

		// Attach attaches a file to the message
		//
		// fileName:	The source file name
		// resType:		Determines if this is a file attachment or an inline object
		// displayName: Name that appears in the attachment section. Applies to attachment type
		// fileId:		Content id to reference in the mail code. Applies to inline type
		Attach(fileName string, resType ResourceType, displayName, fileId string) error

		// ClearAttachments clears file attachments
		ClearAttachments()

		// ClearInlines clears inline files
		ClearInlines()

		// ClearRecipients clears recipients
		ClearRecipients()

		// Encodes the plain text message body to base64 for safe JSON value
		SetMsg(string) error

		// Decodes the encoded message body from base64
		GetMsg() (string, error)
	}
)
