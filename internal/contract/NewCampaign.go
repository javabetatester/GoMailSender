package contract

//DTO
type NewCampaign struct {
	Name    string   `validate:"required"`
	Content string   `validate:"required`
	Emails  []string `validate:"required"`
}
