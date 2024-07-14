package config

import (
	"log"

	"github.com/leetcode-golang-classroom/golang-finance-api/internal/util"
	"github.com/spf13/viper"
)

type Config struct {
	Port                 int64  `mapstructure:"PORT"`
	DbURL                string `mapstructure:"POSTGRES_URL"`
	DefaultAdminPassword string `mapstructure:"DEFAULT_ADMIN_PASSWORD"`
	JWTSignSecret        string `mapstructure:"JWT_SIGN_SECRET"`
}

var AppConfig *Config

func init() {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AutomaticEnv()
	util.FailOnError(v.BindEnv("PORT"), "failed to bind PORT")
	util.FailOnError(v.BindEnv("POSTGRES_URL"), "failed to bind POSTGRES_URL")
	util.FailOnError(v.BindEnv("DEFAULT_ADMIN_PASSWORD"), "failed to bind DEFAULT_ADMIN_PASSWORD")
	util.FailOnError(v.BindEnv("JWT_SIGN_SECRET"), "failed to bind JWT_SIGN_SECRET")
	err := v.ReadInConfig()
	if err != nil {
		log.Println("Load from environment variable")
	}
	err = v.Unmarshal(&AppConfig)
	if err != nil {
		util.FailOnError(err, "Failed to read enivronment")
	}
}
