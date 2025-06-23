package config

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Host string
		Port uint16
	}
	Database struct {
		Postgres struct {
			Host     string
			Port     uint16
			Dbname   string
			User     string
			Password string
			Ssl      string
		}
		Mysql struct {
			Host     string
			Port     uint16
			Dbname   string
			User     string
			Password string
		}
	}
}

func Init() (*Config, error) {
	var env string

	flag.StringVar(&env, "env", "production", "Application environment (default: production)")
	flag.Parse()

	godotenv.Load("./config/" + env + "/.env")

	v := viper.NewWithOptions(
		viper.KeyDelimiter("_"),
	)

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config/" + env + "/")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
		return nil, err
	}

	conf := &Config{}

	if err := v.UnmarshalExact(conf); err != nil {
		log.Fatalf("Error unmarshal config: %s", err)
		return nil, err
	}

	return conf, nil
}
