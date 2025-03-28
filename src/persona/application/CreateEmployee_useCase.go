package application

import "api_recu_corte1/src/persona/domain"

type CreatePerson struct{
	db domain.IPerson
}

func NewCreatePerson (db domain.IPerson) *CreatePerson {
	return &CreatePerson{db: db}
}

func (cp *CreatePerson) Run (person domain.Person) (uint, error) {
	return cp.db.Save(person)
}