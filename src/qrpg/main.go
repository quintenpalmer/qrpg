package qrpg

func Main() {
	model := NewModel()
	model.zone.PersonJoin(NewPerson())
	ListenForever("8000",model)
}
