package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port     string
	BasePath string
	Path     string
	TimeZone string
}
type Thinker_conf struct {
	Endpoint        string
	Email           string
	Password        string
	InputSourceType string
}

type PostgresData struct {
	Host        string
	Port        int
	Username    string
	Password    string
	Dbname      string
	DbSchema    string
	SSLMode     string
	PathMigrate string
}

var PostgresDataConfig *PostgresData
var ConfigData *Config
var ThinkerConfig *Thinker_conf

func LoadConfig() {
	ConfigData = &Config{
		Port:     viper.GetString(`Interface.Http.Port`),
		BasePath: viper.GetString(`Interface.Http.BasePath`),
		Path:     viper.GetString(`Generate.Route.Path`),
		TimeZone: viper.GetString(`System.TimeZone`),
	}

	PostgresDataConfig = &PostgresData{
		Host:        viper.GetString(`Database.Host`),
		Port:        viper.GetInt(`Database.Port`),
		Username:    viper.GetString(`Database.Username`),
		Password:    viper.GetString(`Database.Password`),
		Dbname:      viper.GetString(`Database.DatabaseName`),
		DbSchema:    viper.GetString(`Database.DatabaseSchema`),
		SSLMode:     viper.GetString(`Database.SSLMode`),
		PathMigrate: viper.GetString(`Database.PathMigrate`),
	}

	ThinkerConfig = &Thinker_conf{
		Endpoint:        viper.GetString("Thinker.Endpoint"),
		Email:           viper.GetString("Thinker.Email"),
		Password:        viper.GetString("Thinker.Password"),
		InputSourceType: viper.GetString("Thinker.InputSourceType"),
	}
}
