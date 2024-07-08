package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"pg-go-clean-architecture/conn"
	"pg-go-clean-architecture/migration"
	repo "pg-go-clean-architecture/repositories/club"
	"pg-go-clean-architecture/server"
	service "pg-go-clean-architecture/services/club"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: serve,
}

func serve(cmd *cobra.Command, args []string) {
	db := conn.NewDBClient()

	err := migration.Migrate(db.DB)
	if err != nil {
		println("migration error")
		return
	}

	// repos
	clubRepo := repo.NewClubRepository(db.DB)
	// services
	clubService := service.NewFootballClubService(clubRepo)
	println(clubService)
	//err = clubService.Create(domain.FootballClub{
	//	ClubName: "Chelsea",
	//	Sponsors: []string{"nike", "trivago"},
	//	Id:       1,
	//})
	//if err != nil {
	//	println("error in insert =>")
	//}
	//res, err := clubService.Get(domain.FootballClubPaginatedRequest{
	//	PageSize:   10,
	//	PageNumber: 1,
	//})
	//if err != nil {
	//	println("error in service =>", err.Error())
	//	return
	//}
	//str, err := json.Marshal(res)
	//if err != nil {
	//	println("err=>", err.Error())
	//}
	//println("res=>", string(str))

	// controllers

	var echo_ = echo.New()
	var server = server.New(echo_)
	server.Start()

}
