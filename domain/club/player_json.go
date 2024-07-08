package club

type Player struct {
	ID   int64         `json:"id"`
	Name string        `json:"name"`
	Club *FootballClub `json:"club" pg:"rel:has-one"` // one to one
}
