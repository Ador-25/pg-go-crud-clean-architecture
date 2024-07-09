package club

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"pg-go-clean-architecture/domain/club"
	"time"
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
		return fmt.Errorf("could not insert", err.Error())
	}
	println("insert success", result.RowsAffected())
	return nil
}

func (r Repository) Get(request club.FootballClubPaginatedRequest) ([]club.FootballClub, error) {
	var clubs []club.FootballClub
	err := r.
		dbClient.
		Model(&clubs).
		Where("is_active = ?", request.IsActive).
		Offset((request.PageNumber - 1) * request.PageSize).
		Limit(request.PageSize).
		Select()

	// log on error
	if err != nil {
		println("could not get data")
	}
	return clubs, err
}

func (r Repository) GetById(id int64) (club.FootballClub, error) {
	var fc club.FootballClub
	err := r.
		dbClient.
		Model(&fc).
		Where("id = ? ", id).
		Where("is_active = ?", true).
		Select()

	if err != nil {
		println("could not get club by id:", id, err.Error())
		return club.FootballClub{}, err
	}
	return fc, nil
}

func (r Repository) Delete(id int64) error {
	_, err := r.dbClient.Model(&club.FootballClub{}).
		Set("is_active = ?", false).
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Update()

	if err != nil {
		println("could not set club inactive:", id, err.Error())
		return err
	}
	return nil
}

func (r Repository) Update(fc club.FootballClub) error {
	_, err := r.dbClient.Model(&fc).
		Column("club_name", "sponsors", "updated_at").
		WherePK().
		Update()

	if err != nil {
		println("could not update club:", fc.Id, err.Error())
		return err
	}
	return nil
}
