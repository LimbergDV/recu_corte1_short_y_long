package validators

import (
	"api_recu_corte1/src/persona/domain"
	"errors"
)


func CheckPerson (person domain.Person) error {

	if person.Id < 0{
		return errors.New("El id tiene que ser mayor a 0")
	}

	if person.Name == ""{
		return errors.New("No puedes poner un nombre vacio")
	}

	if person.Age < 0 {
		return errors.New("No puedes poner una edad menor a 0")
	}


	return nil
}