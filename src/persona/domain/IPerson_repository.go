package domain

type IPerson interface {
	Save(person Person) (uint, error)
	GetnewPersonIsAdded() (bool, error)
	CountGender(bool)(int,error)
}