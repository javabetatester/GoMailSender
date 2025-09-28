package campaign

import (
	"GoMailSender/internal/contract"
	"GoMailSender/internal/internalErrors"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
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

func (r *repositoryMock) Delete(id string) error {
	args := r.Called(id)
	return args.Error(0)
}

func (r *repositoryMock) Cancel(id string) error {
	args := r.Called(id)
	return args.Error(0)
}

var (
	testCampaign = contract.NewCampaign{
		Name:    "Test Campaign",
		Content: "Test Content",
		Emails:  []string{"test@example.com"},
	}

	testCampaignEntity = Campaign{
		ID:        "test-id-123",
		Name:      "Test Campaign",
		Content:   "Test Content",
		Status:    Pending,
		CreatedOn: time.Now(),
		Contacts: []Contact{
			{
				ID:         "contact-1",
				Email:      "test@example.com",
				CampaignId: "test-id-123",
			},
		},
	}
)

func Test_Create_Campaign(t *testing.T) {
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.AnythingOfType("*campaign.Campaign")).Return(nil)

	service := Service{Repository: repositoryMock}

	id, err := service.Create(testCampaign)

	assert.NotEmpty(t, id)
	assert.NoError(t, err)
	repositoryMock.AssertExpectations(t)
}

func Test_SaveCampaign(t *testing.T) {
	repositoryMock := new(repositoryMock)

	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		return campaign.Name == testCampaign.Name &&
			campaign.Content == testCampaign.Content &&
			len(campaign.Contacts) == len(testCampaign.Emails)
	})).Return(nil)

	service := Service{Repository: repositoryMock}
	_, err := service.Create(testCampaign)

	assert.NoError(t, err)
	repositoryMock.AssertExpectations(t)
}

// Testes para Delete
func Test_Delete_Campaign_Success(t *testing.T) {
	repositoryMock := new(repositoryMock)

	// Mock GetById - campanha existe
	repositoryMock.On("GetById", "test-id-123").Return(testCampaignEntity, nil)

	// Mock Delete - deleção bem-sucedida
	repositoryMock.On("Delete", "test-id-123").Return(nil)

	service := Service{Repository: repositoryMock}

	err := service.Delete("test-id-123")

	assert.NoError(t, err)
	repositoryMock.AssertExpectations(t)
}

func Test_Delete_Campaign_NotFound(t *testing.T) {
	repositoryMock := new(repositoryMock)

	// Mock GetById - campanha não existe
	repositoryMock.On("GetById", "invalid-id").Return(Campaign{}, gorm.ErrRecordNotFound)

	service := Service{Repository: repositoryMock}

	err := service.Delete("invalid-id")

	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
	repositoryMock.AssertExpectations(t)
}

func Test_Delete_Campaign_DeleteError(t *testing.T) {
	repositoryMock := new(repositoryMock)

	// Mock GetById - campanha existe
	repositoryMock.On("GetById", "test-id-123").Return(testCampaignEntity, nil)

	// Mock Delete - erro na deleção
	deleteError := errors.New("database error")
	repositoryMock.On("Delete", "test-id-123").Return(deleteError)

	service := Service{Repository: repositoryMock}

	err := service.Delete("test-id-123")

	assert.Error(t, err)
	assert.Equal(t, internalErrors.ErrInternalServer, err)
	repositoryMock.AssertExpectations(t)
}
