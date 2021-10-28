package application

import (
	"api/internal/domain"
	"api/pkg/log"
	"fmt"
	"time"
)

type IPersonRepo interface {
	All() []*domain.Person
	Get(ID uint) (*domain.Person, error)
	Create(p *domain.Person) (*domain.Person, error)
	Update(p *domain.Person) (*domain.Person, error)
	Delete(ID uint) error
}

type IPersonService interface {
	GetPeople() []*PersonViewModel
}

type PersonService struct {
	logger     log.IApiLogger
	personRepo IPersonRepo
}

func NewPersonService(logger log.IApiLogger, personRepo IPersonRepo) *PersonService {
	return &PersonService{logger, personRepo}
}

func (ps *PersonService) GetPeople() []*PersonViewModel {
	people := ps.personRepo.All()
	var vm []*PersonViewModel

	for _, person := range people {
		vm = append(vm, &PersonViewModel{
			ID:       fmt.Sprint(person.ID),
			First:    person.FirstName,
			Last:     person.LastName,
			Birthday: person.Birthday.Format("2006-01-02"),
			IsRecent: person.CreatedAt.After(
				time.Now().Add(time.Hour * -2),
			),
		})
	}

	return vm
}

type PersonViewModel struct {
	ID       string `json:"id"`
	First    string `json:"first"`
	Last     string `json:"last"`
	Birthday string `json:"birthday"`
	IsRecent bool   `json:"isRecent"`
}
