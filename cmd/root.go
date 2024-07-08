package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"pg-go-clean-architecture/config"
	"pg-go-clean-architecture/conn"
)

var (
	RootCmd = &cobra.Command{
		Use: "demo",
	}
)

func init() {
	RootCmd.AddCommand(serveCmd)
	//RootCmd.AddCommand(workerCmd)
}

func Execute() {
	// connect to clients
	config.LoadConfig()
	conn.ConnectDB()

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
