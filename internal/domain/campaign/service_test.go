package campaign

import (
	"GoMailSender/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *repositoryMock) Get() ([]Campaign, error) {
	args := r.Called()
	return args.Get(0).([]Campaign), args.Error(1)
}

func (r *repositoryMock) GetById(id string) (Campaign, error) {
	args := r.Called(id)
	return args.Get(0).(Campaign), args.Error(1)
}

var (
	testCampaign = contract.NewCampaign{
		Name:    "Test Campaign",
		Content: "Test Content",
		Emails:  []string{"test@example.com"},
	}
)

func Test_Create_Campaign(t *testing.T) {
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.AnythingOfType("*campaign.Campaign")).Return(nil)
	
	service := Service{Repository: repositoryMock}
	
	id, err := service.Create(testCampaign)

	assert.NotNil(t, id)
	assert.Nil(t, err)
	repositoryMock.AssertExpectations(t)
}

func Test_SaveCampaign(t *testing.T) {
	repositoryMock := new(repositoryMock)

	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != testCampaign.Name {
			return false
		} else if campaign.Content != testCampaign.Content {
			return false
		} else if len(campaign.Contacts) != len(testCampaign.Emails) {
			return false
		}
		return true
	})).Return(nil)
	
	service := Service{Repository: repositoryMock}
	service.Create(testCampaign)
	repositoryMock.AssertExpectations(t)
}
