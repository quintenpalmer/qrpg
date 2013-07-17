package qrpg

import (
	"strings"
)

func Respond(request string, model *Model) (string, error) {
	requests := strings.Split(request,"\n")
	raw := requests[0]
	if raw == "view" {
	} else if raw == "move" {
		dir := requests[1]
		if dir == "up" {
			model.zone.people[0].loc.y -= 1
		} else if dir == "down" {
			model.zone.people[0].loc.y += 1
		} else if dir == "left" {
			model.zone.people[0].loc.x -= 1
		} else if dir == "right" {
			model.zone.people[0].loc.x += 1
		}
	}
	resp, err := model.zone.Serialize()
	if err != nil {
		return "", err
	} else {
		return resp.toString()
	}
}
