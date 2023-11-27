package config

type JWT struct {
	SignKey     string `mapstructure:"sign-key" json:"sign-key" yaml:"sign-key"`
	ExpiredTime string `mapstructure:"expired-time" json:"expired-time" yaml:"expired-time"`
	BufferTime  string `mapstructure:"buffer-time" json:"buffer-time" yaml:"buffer-time"`
	IssUer      string `mapstructure:"issuser" json:"issuer" yaml:"issuer"`
}
