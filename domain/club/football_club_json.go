package club

import "time"

type FootballClub struct {
	Id        int64     `json:"id"`
	ClubName  string    `json:"club_name"`
	Sponsors  []string  `json:"sponsors" pg:",array"`
	Players   []Player  `json:"players" pg:"rel:has-many"`
	CreatedAt time.Time `json:"created_at" pg:"default:now()"`
	UpdatedAt time.Time `json:"updated_at" pg:"default:now()"`
	IsActive  bool      `json:"is_active" pg:"default:true"`
}

type FootballClubPaginatedRequest struct {
	PageNumber int
	PageSize   int
	IsActive   *bool
}
