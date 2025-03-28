package domain

type Person struct {
	Id int32
	Name string
	Age int
	Gender bool
}

func NewPerson(Id int32, Name string, Age int, Gender bool) *Person{
	return &Person {Id: 1, Name: Name, Age:Age, Gender: Gender}
}
