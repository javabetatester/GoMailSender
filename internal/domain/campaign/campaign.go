package campaign

import (
	"errors"

	"time"
	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"required,email"`
}
type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=100"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=10,max=1024"`
	Contacts  []Contact `validate:"min=1"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))

	for index, value := range emails {
		contacts[index].Email = value
	}

	if name == "" || content == "" || len(emails) == 0 {
		return nil, errors.New("name is required")
	} else if content == "" {
		return nil, errors.New("content is required")
	} else if len(emails) == 0 {
		return nil, errors.New("at least one email is required")
	}

	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}, nil
}
