package domain

type IPerson interface {
	Save(person Person) (uint, error)
	
}