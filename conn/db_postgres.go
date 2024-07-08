package conn

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
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
		println("could not connect client")
		return
	}
	println("client connected")
}

func CloseDB() {
	postgres_client.DB.Close()
}
