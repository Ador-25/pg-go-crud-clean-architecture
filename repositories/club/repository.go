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
	result, err := r.dbClient.
		Model(&fc).
		Insert()
	// log on error
	if err != nil {
		println("could not insert", err.Error())
		return err
	}
	println("insert success", result.RowsAffected())
	return nil
}

func (r Repository) Get(request club.FootballClubPaginatedRequest) ([]club.FootballClub, error) {
	var clubs []club.FootballClub
	err := r.
		dbClient.
		Model(&clubs).
		Offset((request.PageNumber - 1) * request.PageSize).
		Limit(request.PageSize).
		Select()

	// log on error
	if err != nil {
		println("could not get data")
	}
	return clubs, err
}
