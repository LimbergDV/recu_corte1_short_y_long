package domain

type Person struct {
	Id int32
	Name string
	Sex bool
	Gender string
}

func NewPerson(Id int32, Name string, Sex bool, Gender string) *Person{
	return &Person {Id: 1, Name: Name, Sex: Sex, Gender: Gender}
}
