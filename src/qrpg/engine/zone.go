package engine

type Zone struct {
	people []*Person
}

func NewZone() *Zone {
	zone := new(Zone)
	zone.people = make([]*Person,0)
	return zone
}

func (zone *Zone) GetPersonFromName(name string) (*Person, error) {
	for _, person := range zone.people {
		if person.name == name {
			return person, nil
		}
	}
	return nil, NewQrpgError("Player not in Zone")
}

func (zone *Zone) PersonJoin(person *Person) error {
	zone.people = append(zone.people,person)
	return nil
}

func (zone *Zone) PersonInZone(person *Person) (bool, int) {
	for i, zonePerson := range zone.people {
		if person == zonePerson {
			return true, i
		}
	}
	return false, -1
}

func (zone *Zone) PersonLeave(person *Person) error {
	inZone, index := zone.PersonInZone(person)
	if inZone {
		zone.people = append(zone.people[:index], zone.people[index:]...)
		return nil
	} else {
		return NewQrpgError("Person not in zone!")
	}
}
