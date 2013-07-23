package engine

type Battle struct {
	people []*Person
}

func NewBattle(player1 *Person, player2 *Person) *Battle {
	battle := new(Battle)
	battle.people = make([]*Person,0,2)
	battle.people = append(battle.people,player1)
	battle.people = append(battle.people,player2)
	return battle
}
