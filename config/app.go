package config

type Configuration struct {
	Chain    Chain    `mapstructure:"Chain" json:"chain" yaml:"chain"`
	Database Database `mapstructure:"database" json:"database" yaml:"database"`
	Persist  Persist  `mapstructure:"persist" json:"persist" yaml:"persist"`
}
