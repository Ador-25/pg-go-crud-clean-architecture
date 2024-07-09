package club

type FootballClubRepository interface {
	Create(club FootballClub) error
	Get(request FootballClubPaginatedRequest) ([]FootballClub, error)
	Update(club FootballClub) error
	Delete(id int64) error
	GetById(id int64) (FootballClub, error)
}

type FootballClubService interface {
	Create(club FootballClub) error
	Get(request FootballClubPaginatedRequest) ([]FootballClub, error)
	Update(club FootballClub) error
	Delete(id int64) error
	GetById(id int64) (FootballClub, error)
}
