package notifykit

import (
	"time"
)

// EmailFormat
const (
	MAIL_FORMAT_PLAIN MailFormat = "plain"
	MAIL_FORMAT_HTML  MailFormat = "html"
)

// Importance
const (
	PRIORITY_NORMAL MailPriority = 0
	PRIORITY_LOW    MailPriority = 1
	PRIORITY_HIGH   MailPriority = 2
)

type (
	MailFormat   string
	MailPriority uint8

	MailParticipant struct {
		Name    string `json:"name,omitempty"`
		Address string `json:"address,omitempty"`
		Role    *Role  `json:"role,omitempty"`
	}
	MailRecipient struct {
		MailParticipant
		Key     string `json:"key,omitempty"`
		MailKey string `json:"mail_key,omitempty"`
	}
	MailAttachment struct {
		Key                string       `json:"key,omitempty"`           // Attachment key
		MailKey            string       `json:"mail_key,omitempty"`      // Mail key
		ResourceType       ResourceType `json:"type,omitempty"`          // Resource type: Attachment or Inline Resource
		AttachmentID       string       `json:"attachment_id,omitempty"` // Attachment id for inline resource
		ResourceID         string       `json:"resource_id,omitempty"`   // ID for blob
		FileName           string       `json:"file_name,omitempty"`     // File name
		DisplayName        string       `json:"name,omitempty"`          // Name of the attachment to display
		UploadBaseFileName string
	}
	Mail struct {
		Key                 string           `json:"key,omitempty"`
		ApplicationID       string           `json:"app_id,omitempty"`
		From                MailParticipant  `json:"from"`
		Sender              *MailParticipant `json:"sender,omitempty"`
		ReplyTo             *MailParticipant `json:"reply_to,omitempty"`
		Subject             string           `json:"subject,omitempty"`
		Body                string           `json:"body,omitempty"`
		Format              MailFormat       `json:"format,omitempty"`
		Priority            MailPriority     `json:"priority,omitempty"`
		NoFooter            bool             `json:"no_footer,omitempty"`
		NoSubjectID         bool             `json:"no_subject_id,omitempty"`
		DelaySend           bool             `json:"delay_send,omitempty"`
		DelaySendOn         *time.Time       `json:"delay_send_on,omitempty"`
		ExpectedAttachments int              `json:"expected_attachments,omitempty"`
		ExpectedInlines     int              `json:"expected_inlines,omitempty"`
		Recipients          []MailRecipient  `json:"recipients,omitempty"`
		Inlines             []MailAttachment `json:"inlines,omitempty"`
		Attachments         []MailAttachment `json:"attachments,omitempty"`
	}
)
