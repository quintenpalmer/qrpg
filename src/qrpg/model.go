package qrpg

type Model struct {
	zone *Zone
}

func NewModel() *Model {
	model := new(Model)
	model.zone = NewZone()
	return model
}
