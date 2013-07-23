package engine

func Main() {
	model := NewModel()
	model.zone.PersonJoin(NewPerson("Quin",0,0))
	model.zone.PersonJoin(NewPerson("Cloud",0,1))
	ListenForever("8000",model)
}
