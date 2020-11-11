package pkg

import (
	"github.com/spf13/viper"
)

const prefix = "postgres."

type Config struct {
	Host              string `mapstructure:"host"`
	Port              int    `mapstructure:"port"`
	Username          string `mapstructure:"username"`
	Password          string `mapstructure:"password"`
	Database          string `mapstructure:"database"`
	ConnectionOptions string `mapstructure:"connection-options"`
}

func (c *Config) InitFromViper(v *viper.Viper) error {
	v.SetDefault(prefix+"port", 5432)
	v.SetDefault(prefix+"database", "jaeger")

	c.Host = v.GetString(prefix + "host")
	c.Port = v.GetInt(prefix + "port")
	c.Username = v.GetString(prefix + "username")
	c.Password = v.GetString(prefix + "password")
	c.Database = v.GetString(prefix + "database")
	c.ConnectionOptions = v.GetString(prefix + "connection-options")
	return nil
}
