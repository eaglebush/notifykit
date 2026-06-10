package notifykit

import (
	"encoding/base64"
	"fmt"
	"os"
	"regexp"
	"strings"
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

// NewEmail creates an initialized mail queue
func NewMail(mo ...MailOption) (*Mail, error) {
	m := Mail{
		Recipients:  make([]MailRecipient, 0),
		Attachments: make([]MailAttachment, 0),
		Inlines:     make([]MailAttachment, 0),
	}
	for _, opt := range mo {
		opt(&m)
	}
	if m.ApplicationID == "" {
		return nil, ErrMailHasNoApplicationId
	}
	return &m, nil
}

// SetMsg encodes the plain text mail body for safe JSON value
func (obj *Mail) SetMsg(plainText string) error {
	if plainText == "" {
		return ErrMailMessageNotProvided
	}
	obj.Body = base64.StdEncoding.EncodeToString([]byte(plainText))
	return nil
}

// GetMsg decodes the encoded mail body
func (obj *Mail) GetMsg() (string, error) {
	if obj.Body == "" {
		return "", nil
	}
	body, err := base64.StdEncoding.DecodeString(obj.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// AddRecipient adds recipient to the mail message
func (obj *Mail) AddRecipient(name string, address string, role Role) error {
	if name == "" {
		return ErrMailRecipientNameNotProvided
	}
	if address == "" {
		return ErrMailRecipientAddressNotProvided
	}
	if err := validateEmail(&address); err != nil {
		return fmt.Errorf("address %s", err)
	}
	for _, v := range obj.Recipients {
		if strings.EqualFold(address, v.Address) && *v.Role == role {
			return nil
		}
	}
	obj.Recipients =
		append(
			obj.Recipients,
			MailRecipient{
				MailParticipant: MailParticipant{
					Name:    name,
					Address: address,
					Role:    &role,
				},
			})

	return nil
}

// Attach attaches a file to the message
func (obj *Mail) Attach(fileName string, resType ResourceType, displayName, fileId string) error {
	if fileName == "" {
		return ErrMailFileIsNotProvided
	}

	// Check if exists
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			return ErrMailFileDoesNotExist
		}
	}
	for _, v := range obj.Attachments {
		if strings.EqualFold(fileName, v.FileName) {
			return nil
		}
	}

	if resType == ATTACHMENT {
		obj.Attachments =
			append(obj.Attachments,
				MailAttachment{
					FileName:     fileName,
					DisplayName:  displayName,
					ResourceType: resType,
					AttachmentID: fileId,
				})
		obj.ExpectedAttachments = len(obj.Attachments)
	}
	if resType == INLINE {
		obj.Inlines =
			append(obj.Inlines,
				MailAttachment{
					FileName:     fileName,
					DisplayName:  displayName,
					ResourceType: resType,
					AttachmentID: fileId,
				})
		obj.ExpectedInlines = len(obj.Inlines)
	}
	return nil
}

// ClearAttachments clears file attachments
func (obj *Mail) ClearAttachments() {
	atts := make([]MailAttachment, 0)
	for _, att := range obj.Attachments {
		if att.ResourceType != ATTACHMENT {
			atts = append(atts, att)
		}
	}
	obj.Attachments = atts
	obj.ExpectedAttachments = 0
}

// ClearInlines clears inline attachments
func (obj *Mail) ClearInlines() {
	ress := make([]MailAttachment, 0)
	for _, att := range obj.Attachments {
		if att.ResourceType != INLINE {
			ress = append(ress, att)
		}
	}
	obj.Inlines = ress
	obj.ExpectedInlines = 0
}

// ClearRecipients clears recipients
func (obj *Mail) ClearRecipients() {
	obj.Recipients = make([]MailRecipient, 0)
}

func validateEmail(email *string) error {
	if email == nil || *email == "" {
		return fmt.Errorf("is empty")
	}
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9]" +
		"(?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(*email) {
		return fmt.Errorf("is an invalid email address")
	}
	return nil
}
