package config

type DBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

type AppConfig struct {
	Name         string
	Port         string
	AppKeyHeader string
	AppKey       string
}
type Config struct {
	App *AppConfig
	DB  *DBConfig
}

var config Config

func DB() *DBConfig {
	return config.DB
}
func App() *AppConfig {
	return config.App
}
func setDefaultConfig() {
	config.App = &AppConfig{
		Name:         "pg-crud",
		Port:         "5555",
		AppKeyHeader: "demo-app-key",
		AppKey:       "fb60e941d7c48db4810d2a4282732acf",
	}
	config.DB = &DBConfig{
		Host:     "127.0.0.1",
		Port:     5433,
		Username: "postgres",
		Password: "postgres",
	}
}
func LoadConfig() {
	setDefaultConfig()
}
