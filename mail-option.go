package notifykit

import (
	"time"
)

type MailOption func(*Mail) error

// ApplicationID sets the application path
func ApplicationID(id string) MailOption {
	return func(m *Mail) error {
		m.ApplicationID = id
		return nil
	}
}

// Subject sets the subject
func Subject(subject string) MailOption {
	return func(m *Mail) error {
		m.Subject = subject
		return nil
	}
}

// Format sets the e-mail format
func Format(ft MailFormat) MailOption {
	return func(m *Mail) error {
		m.Format = ft
		return nil
	}
}

// Priority sets the e-mail importance
func Priority(pri MailPriority) MailOption {
	return func(m *Mail) error {
		m.Priority = pri
		return nil
	}
}

// From sets the e-mail author name and address
func From(name, address string) MailOption {
	return func(m *Mail) error {
		f := FROM
		m.From = MailParticipant{
			Name:    name,
			Address: address,
			Role:    &f,
		}
		return nil
	}
}

// Origin sets the e-mail sender and address
func Origin(name, address string) MailOption {
	return func(m *Mail) error {
		s := SENDER
		m.Sender = &MailParticipant{
			Name:    name,
			Address: address,
			Role:    &s,
		}
		return nil
	}
}

// ReplyTo sets the e-mail name and address of the recipient to reply to
func ReplyTo(name, address string) MailOption {
	return func(m *Mail) error {
		r := REPLYTO
		m.ReplyTo = &MailParticipant{
			Name:    name,
			Address: address,
			Role:    &r,
		}
		return nil
	}
}

// To sets the e-mail name and address of the recipient
func To(name, address string) MailOption {
	return func(m *Mail) error {
		m.AddRecipient(name, address, TO)
		return nil
	}
}

// Cc sets the e-mail name and address of the copy furnish recipient
func Cc(name, address string) MailOption {
	return func(m *Mail) error {
		m.AddRecipient(name, address, CC)
		return nil
	}
}

// Bcc sets the e-mail name and address of the recipient as blind carbon copy
func Bcc(name, address string) MailOption {
	return func(m *Mail) error {
		m.AddRecipient(name, address, BCC)
		return nil
	}
}

// NoFooter sets whether the e-mail should have a footer
func NoFooter(value bool) MailOption {
	return func(m *Mail) error {
		m.NoFooter = value
		return nil
	}
}

// NoSubjectID sets whether the e-mail should have a subject id
func NoSubjectID(value bool) MailOption {
	return func(m *Mail) error {
		m.NoSubjectID = value
		return nil
	}
}

// DelaySend sets whether the e-mail should be sent later
func DelaySend(sendBy *time.Time) MailOption {
	return func(m *Mail) error {
		m.DelaySend = true
		m.DelaySendOn = sendBy
		return nil
	}
}

// Body sets the message of the mail.
func Body(msg string) MailOption {
	return func(m *Mail) error {
		m.SetMsg(msg)
		return nil
	}
}
