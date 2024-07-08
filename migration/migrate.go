package migration

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"pg-go-clean-architecture/domain/club"
)

func Migrate(db *pg.DB) error {
	models := []interface{}{
		(*club.FootballClub)(nil),
		(*club.Player)(nil),
	}
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			println("error creating table", err.Error())
			return err
		}
	}
	println("migration done")
	return nil
}
