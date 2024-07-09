package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"pg-go-clean-architecture/conn"
	"pg-go-clean-architecture/controllers"
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

	// dependency injection
	clubRepo := repo.NewClubRepository(db.DB)
	clubService := service.NewFootballClubService(clubRepo)
	clubController := controllers.NewClubController(clubService)

	// endpoints

	var e = echo.New()
	e.GET("/clubs", clubController.GetClubs)
	e.GET("/clubs/:id", clubController.GetClubByID)
	e.POST("/clubs", clubController.CreateClub)
	e.PUT("/clubs/:id", clubController.UpdateClub)
	e.DELETE("/clubs/:id", clubController.DeleteClub)
	var server = server.New(e)
	server.Start()

}
