package club

import "time"

type Player struct {
	ID        int64         `json:"id"`
	Name      string        `json:"name"`
	Club      *FootballClub `json:"club" pg:"rel:has-one"`
	ClubID    int64         `json:"club_id"`
	CreatedAt time.Time     `json:"created_at" pg:"default:now()"`
	UpdatedAt time.Time     `json:"updated_at" pg:"default:now()"`
	IsActive  bool          `json:"is_active" pg:"default:true"`
}
