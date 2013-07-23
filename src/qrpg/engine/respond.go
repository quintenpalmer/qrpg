package engine

import (
	"strings"
	"fmt"
)

func Respond(request string, model *Model) (string, error) {
	fmt.Println(request)
	requests := strings.Split(request,",")
	raw := requests[0]
	what := requests[1]
	if raw == "zone" {
		active_player := false
		if what == "move" {
			who := requests[2]
			dir := requests[3]
			active_player = true
			model.zone.PersonMove(who,dir)
		} else if what == "view" {

		}
		model.zone.startBattle()
		if active_player {
			who := requests[2]
			person, err := model.zone.GetPersonFromName(who)
			if err != nil {
				return "", err
			}
			if person.inBattle {
				return respondBattle(model)
			}
		}
		return respondZone(model)
	} else if raw == "battle" {
		//who := requests[2]
		//person, err := model.zone.GetPersonFromName(who)
		if what == "finish" {
			model.zone.endBattles()
			return respondZone(model)
		}
		return respondBattle(model)
	} else {
		return "", NewQrpgError("Not a valid command given")
	}
}

func respondBattle(model *Model) (string, error) {
	return `{ "type" : "battle" }`, nil
}

func respondZone(model *Model) (string, error) {
	resp, err := model.zone.Serialize()
	if err != nil {
		return "", err
	} else {
		return resp.toString()
	}
}
