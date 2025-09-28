package campaign

type Repository interface {
	Save(campaign *Campaign) error
	Get() ([]Campaign, error)
	GetById(id string) (Campaign, error)
	Delete(id string) error
	Cancel(id string) error
}
