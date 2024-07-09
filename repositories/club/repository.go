package club

import (
	"github.com/go-pg/pg/v10"
	"log/slog"
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
	_, err := r.dbClient.
		Model(&fc).
		Column("club_name", "created_at", "sponsors", "is_active", "updated_at").
		Insert()
	// log on error
	if err != nil {
		slog.Error("could not insert", map[string]any{
			"err":  err.Error(),
			"data": fc,
		})
		return err
	}
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
		slog.Error("could not get data", map[string]any{
			"err":          err.Error(),
			"request_data": request,
		})
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
		slog.Error("could not get data", map[string]any{
			"err": err.Error(),
			"id":  id,
		})
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
		slog.Error("could not delete", map[string]any{
			"err": err.Error(),
			"id":  id,
		})
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
		slog.Error("could not update", map[string]any{
			"err":  err.Error(),
			"club": fc,
		})
		return err
	}
	return nil
}
