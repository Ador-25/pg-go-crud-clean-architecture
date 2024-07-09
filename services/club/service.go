package club

import (
	domain "pg-go-clean-architecture/domain/club"
	"time"
)

type FootballClubService struct {
	repository domain.FootballClubRepository
}

func NewFootballClubService(repo domain.FootballClubRepository) *FootballClubService {
	return &FootballClubService{repository: repo}
}
func (s FootballClubService) Get(request domain.FootballClubPaginatedRequest) ([]domain.FootballClub, error) {
	var isActive bool
	// get active data by default
	if request.IsActive == nil {
		isActive = true
	} else {
		isActive = *request.IsActive
	}

	request.IsActive = &isActive
	return s.repository.Get(request)
}

func (s FootballClubService) Create(fc domain.FootballClub) error {
	currentTime := time.Now()
	fc.IsActive = true
	fc.CreatedAt = currentTime
	fc.UpdatedAt = currentTime
	return s.repository.Create(fc)
}
func (s FootballClubService) Update(club domain.FootballClub) error {
	currentTime := time.Now()
	club.UpdatedAt = currentTime
	return s.repository.Update(club)
}
func (s FootballClubService) Delete(id int64) error {
	return s.repository.Delete(id)
}
func (s FootballClubService) GetById(id int64) (domain.FootballClub, error) {
	data, err := s.repository.GetById(id)
	return data, err
}
