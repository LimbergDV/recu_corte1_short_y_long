package application

import "api_recu_corte1/src/persona/domain"



type CountGenderUc struct {
	db domain.IPerson
}

func NewCountGenderUc(db domain.IPerson)*CountGenderUc{
	return &CountGenderUc{db: db}
}

func (useCase *CountGenderUc)Execute(sexo bool)(int,error){
	return useCase.db.CountGender(sexo)
}