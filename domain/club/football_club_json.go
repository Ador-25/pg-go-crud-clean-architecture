package club

type FootballClub struct {
	Id       int64    `json:"id"`
	ClubName string   `json:"club_name"`
	Sponsors []string `json:"sponsors"`                  // array
	Players  []Player `json:"players" pg:"rel:has-many"` // one to many
}

type FootballClubPaginatedRequest struct {
	PageNumber int
	PageSize   int
}
