package infrastructure

import (
	"api/internal/domain"
	"api/pkg/log"
)

type PersonRepo struct {
	logger log.IApiLogger
	db     *SQLDatabase
}

func NewPersonRepo(logger log.IApiLogger, db *SQLDatabase) *PersonRepo {
	return &PersonRepo{logger, db}
}

func (pr *PersonRepo) All() []*domain.Person {
	var people []*domain.Person
	pr.db.Find(&people)

	return people
}

func (pr *PersonRepo) Get(ID uint) (*domain.Person, error) {
	return nil, nil
}

func (pr *PersonRepo) Create(p *domain.Person) (*domain.Person, error) {
	return nil, nil
}

func (pr *PersonRepo) Update(p *domain.Person) (*domain.Person, error) {
	return nil, nil
}

func (pr *PersonRepo) Delete(ID uint) error {
	return nil
}
