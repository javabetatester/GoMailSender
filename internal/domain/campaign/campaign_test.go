package campaign

import "testing"

func TestNewCampaign(t *testing.T) {
	name := "Campaign x"
	content := "Body"
	contacts := []string{"email1@example.com", "email2@example.com"}

	campaign := NewCampaign(name, content, contacts)

	if campaign.ID != "1" {
		t.Errorf("Expected %s, got %s", "1", campaign.ID)
	} else if campaign.Name != name {
		t.Errorf("Expected %s, got %s", name, campaign.Name)
	} else if campaign.Content != content {
		t.Errorf("Expected %s, got %s", content, campaign.Content)
	} else if len(campaign.Contacts) != len(contacts) {
		t.Errorf("Expected %d, got %d", len(contacts), len(campaign.Contacts))
	}
}
