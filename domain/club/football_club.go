package club

type FootballClubRepository interface {
	Create(club FootballClub) error
	Get(request FootballClubPaginatedRequest) ([]FootballClub, error)
}
