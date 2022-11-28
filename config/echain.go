package config

type Chain struct {
	Url          string `mapstructure:"url" json:"url" yaml:"url"`
	Method       string `mapstructure:"method" json:"method" yaml:"method"`
	BlockAddress string `mapstructure:"block_address" json:"block_address" yaml:"block_address"`
}
