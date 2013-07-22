package engine

type Person struct {
	loc *Loc
}

func NewPerson() *Person {
	person := new(Person)
	person.loc = NewLoc(0,0)
	return person
}
