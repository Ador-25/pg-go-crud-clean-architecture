package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"pg-go-clean-architecture/conn"
	"pg-go-clean-architecture/migration"
	"pg-go-clean-architecture/server"
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
	println(db)

	// repos

	// services

	// controllers

	var echo_ = echo.New()
	var server = server.New(echo_)
	server.Start()

}
