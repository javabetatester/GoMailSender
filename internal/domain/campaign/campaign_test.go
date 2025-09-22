package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign x"
	content  = "Body"
	contacts = []string{"email1@example.com", "email2@example.com"}
)

// TDD com Testes nativos do Golang
func Test_NewCampaign_Go(t *testing.T) {

	campaign, _ := NewCampaign(name, content, contacts)

	if campaign.ID == "" {
		t.Errorf("Expected ID to be set, got empty string")
	} else if campaign.Name != name {
		t.Errorf("Expected %s, got %s", name, campaign.Name)
	} else if campaign.Content != content {
		t.Errorf("Expected %s, got %s", content, campaign.Content)
	} else if len(campaign.Contacts) != len(contacts) {
		t.Errorf("Expected %d, got %d", len(contacts), len(campaign.Contacts))
	}
}

// TDD com Testify
func Test_NewCampaign_CreateCampaign(t *testing.T) {

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotEmpty(t, campaign.ID)
	assert.Equal(t, name, campaign.Name)
	assert.Equal(t, content, campaign.Content)
	assert.Equal(t, len(contacts), len(campaign.Contacts))
}

func Test_NewCampaign_IDIsNotNull(t *testing.T) {

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(t, campaign.ID)
}

func Test_NewCampaign_CreatedOnIsNotNull(t *testing.T) {

	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Greater(t, campaign.CreatedOn, now)
}

func Test_NewCampaign_HasContent(t *testing.T) {

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Equal(t, content, campaign.Content)
}


//TESTE COM VARIOS CASES
func Test_NewCampaign_MustValidateName(t *testing.T) {

	_, err := NewCampaign("", content, contacts)

	assert.Equal(t, "name is required", err.Error())

	_, err = NewCampaign("oioi", content, contacts)

	assert.Equal(t, "name must be at least 5 characters", err.Error())
}
