package club

import (
	"github.com/go-pg/pg/v10"
	"pg-go-clean-architecture/domain/club"
)

type Repository struct {
	dbClient *pg.DB
}

func NewClubRepository(conn *pg.DB) *Repository {
	return &Repository{
		dbClient: conn,
	}
}

func (r Repository) Create(fc club.FootballClub) error {
	// implement
	return nil
}

func (r Repository) Get(request club.FootballClubPaginatedRequest) ([]club.FootballClub, error) {
	// implement
	return nil, nil
}
