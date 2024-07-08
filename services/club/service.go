package club

import (
	domain "pg-go-clean-architecture/domain/club"
)

type FootballClubService struct {
	repository domain.FootballClubRepository
}

func NewFootballClubService(repo domain.FootballClubRepository) *FootballClubService {
	return &FootballClubService{repository: repo}
}
func (s FootballClubService) Get(request domain.FootballClubPaginatedRequest) ([]domain.FootballClub, error) {
	return s.repository.Get(request)
}

func (s FootballClubService) Create(fc domain.FootballClub) error {
	return s.repository.Create(fc)
}
