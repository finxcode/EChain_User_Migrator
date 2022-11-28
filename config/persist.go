package config

type Persist struct {
	Loc      string `mapstructure:"loc" json:"loc" yaml:"loc"`
	Filename string `mapstructure:"filename" json:"filename" yaml:"filename"`
}
