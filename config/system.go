package config

type System struct {
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`                         // 端口值
	DbType       string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`                // 数据库类型
	UseRedis     bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`          // 使用redis
	UseDatabase  bool   `mapstructure:"use-database" json:"use-database" yaml:"use-database"` // 使用数据库
	LimitCountIP int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP  int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
}
