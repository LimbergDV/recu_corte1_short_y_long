package domain

type IPerson interface {
	Save(person Person) (uint, error)
	NewPersonIsAdded(person Person) (uint, error)
	CountGender(person Person) (uint, error)
}