package engine

type Person struct {
	loc *Loc
	name string
}

func NewPerson(name string, x int, y int) *Person {
	person := new(Person)
	person.name = name
	person.loc = NewLoc(x,y)
	return person
}
