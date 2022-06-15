package config

type Server struct {
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email  Email  `mapstructure:"email" json:"email" yaml:"email"`
	System System `mapstructure:"system" json:"system" yaml:"system"`

	// gorm
	Mysql     Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Pgsql     Pgsql           `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Cockroach Pgsql           `mapstructure:"cockroach" json:"cockroach" yaml:"cockroach"`
	DBList    []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
}
