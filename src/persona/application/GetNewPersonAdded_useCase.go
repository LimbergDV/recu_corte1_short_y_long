package application

import "api_recu_corte1/src/persona/domain"


type GetNewPersonIsAddedUc struct {
	db domain.IPerson
}

func NewGetNewPersonIsAddedUc(db domain.IPerson)*GetNewPersonIsAddedUc{
	return &GetNewPersonIsAddedUc{db: db}
}

func (useCase *GetNewPersonIsAddedUc)Execute()(bool, error){
	return useCase.db.GetnewPersonIsAdded()
}