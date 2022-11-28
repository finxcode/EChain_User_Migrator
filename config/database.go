package config

type Database struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"Port" json:"Port" yaml:"Port"`
	Database string `mapstructure:"database" json:"database" yaml:"database"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
