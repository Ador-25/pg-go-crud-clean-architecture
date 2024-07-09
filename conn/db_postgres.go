package conn

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"log/slog"
	"pg-go-clean-architecture/config"
	"strconv"
)

var postgres_client DBClient

type DBClient struct {
	DB *pg.DB
}

func NewDBClient() *DBClient {
	return &postgres_client
}

func ConnectDB() {
	c := config.DB()
	client := pg.Connect(&pg.Options{
		Database: "postgres",
		User:     c.Username,
		Password: c.Password,
		Addr:     fmt.Sprintf("%s:%s", c.Host, strconv.Itoa(c.Port)),
	})
	postgres_client.DB = client
	err := client.Ping(context.Background())
	if err != nil {
		slog.Error("could not connect to db", map[string]any{
			"error": err.Error(),
		})
		return
	}
	slog.Info("connected to db")
}

func CloseDB() {
	err := postgres_client.DB.Close()
	if err != nil {
		slog.Error("could not close db client", map[string]any{
			"error": err.Error(),
		})
		return
	}
	slog.Info("closed db client")
}
