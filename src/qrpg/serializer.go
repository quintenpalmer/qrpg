package qrpg

import (
	"encoding/json"
)

type jsonMap map[string]interface{}

func (repr jsonMap) toString() (string, error) {
	str, err := json.Marshal(repr)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

func (zone *Zone) Serialize() (jsonMap, error) {
	repr := make(jsonMap)
	people := make([]jsonMap,0,len(zone.people))
	for _, person := range zone.people {
		person_repr, err := person.Serialize()
		if err != nil {
			return nil, NewQrpgError("Bad player data")
		}
		people = append(people,person_repr)
	}
	repr["people"] = people
	return repr, nil
}

func (person *Person) Serialize() (jsonMap, error) {
	repr := make(jsonMap)
	repr["x"] = person.loc.x
	repr["y"] = person.loc.y
	return repr, nil
}
