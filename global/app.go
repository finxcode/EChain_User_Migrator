package global

import (
	"echain-user-updater/config"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	DB          *sqlx.DB
}

var App = new(Application)
