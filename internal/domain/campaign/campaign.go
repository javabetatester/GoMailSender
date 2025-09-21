package campaign

import (
	"errors"

	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string
}
type Campaign struct {
	ID        string
	Name      string
	CreatedOn time.Time
	Content   string
	Contacts  []Contact
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))

	for index, value := range emails {
		contacts[index].Email = value
	}

	if name == "" || content == "" || len(emails) == 0 {
		return nil, errors.New("invalid campaign data")
	}

	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}, nil
}
