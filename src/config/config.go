package config

import (
	"io/ioutil"
	"github.com/BurntSushi/toml"
)

type Conf struct {
	Database    DatabaseConf   `toml:"Database"`
	Server      ServerConf     `toml:"Server"`
}

type DatabaseConf struct {
	SQLString   string    `toml:"MySQLString"`
	RedisAddr   string    `toml:"RedisAddress"`
}

type ServerConf struct {
	Port   string   `toml:"Port"`
}

func (c *Conf) GetConfig() (*Conf, error) {
	confFile, err := ioutil.ReadFile("config.toml")
	if err != nil {
		return c, err
	}
	err = toml.Unmarshal(confFile, c)
	if err != nil {
		return c, err
	}
	return c, nil
}

var C Conf
