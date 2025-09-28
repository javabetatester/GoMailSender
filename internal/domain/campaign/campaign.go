package campaign

import (
	"errors"

	"time"

	"github.com/rs/xid"
)

const (
	Pending  string = "Pending"
	Active   string = "Active"
	Finished string = "Finished"
	Canceled string = "Canceled"
)

type Contact struct {
	ID         string `gorm:"size:50"`
	Email      string `validate:"required,email"`
	CampaignId string `gorm:"size:50"`
}

type Campaign struct {
	ID        string    `validate:"required" gorm:"size:50"`
	Name      string    `validate:"min=5,max=100" gorm:"size:100"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=10,max=1024" gorm:"size:1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string    `validate:"required" gorm:"size:15"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))

	for index, value := range emails {
		contacts[index].Email = value
		contacts[index].ID = xid.New().String()
	}

	// Validações de disparo de erro são necessárias implementadas, todos os cases que disparam um erro de tipo "x"
	if name == "" {
		return nil, errors.New("name is required")
	}
	if len(name) < 5 {
		return nil, errors.New("name must be at least 5 characters")
	}
	if content == "" {
		return nil, errors.New("content is required")
	}
	if len(content) < 10 {
		return nil, errors.New("content must be at least 10 characters")
	}
	if len(emails) == 0 {
		return nil, errors.New("at least one email is required")
	}

	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
		Status:    Pending,
	}, nil
}
