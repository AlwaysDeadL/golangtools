package types

type Animal string
type Person string

type Teacher struct {
	Attribute Person
	Name      string
	age       int
}

func GetTeacher(name string) Teacher {
	return Teacher{
		Attribute: Person("name"),
		Name:      name,
		age:       0,
	}
}
