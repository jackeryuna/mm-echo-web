package conf

import (
	"errors"
	"os"
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/labstack/gommon/log"
)

var (
	Conf config
	DefaultConfigFile string = "conf/conf.toml"
)

type config struct {
	ReleaseMode bool `toml:"release_mode"`
	LogLevel string `toml:"log_level"`

	SessionStore string `toml:"session_store"`
	CacheStore string `toml:"cache_store"`

	App app

	Server server

	DB database `toml:"database"`

	Redis redis

}

type app struct {
	Name string `toml:"name"`
}

type server struct {
	Graceful bool `toml:"graceful"`
	Addr string `toml:"addr"`

	DomainApi string `toml:"domain_api"`
	DomainWeb string `toml:"domain_web"`
	DomainSocket string `toml:"domain_socket"`
}

type database struct {
	Name string `toml:"name"`
	UserName string `toml:"user_name"`
	Pwd string `toml:"pwd"`
	Host string `toml:"host"`
	Port string `toml:"port"`
}

type redis struct {
	Server string `toml:"server"`
	Pwd string `toml:"pwd"`
}

func InitConfig(configFile string) error {
	if configFile == "" {
		configFile = DefaultConfigFile
	}

	Conf = config{
		ReleaseMode : false,
		LogLevel : "DEBUG",
	}

	if _, err := os.Stat(configFile); err != nil {
		return errors.New("config file error:" + err.Error())
	} else {
		log.Info("load config from file:" + configFile)
		configBytes, err := ioutil.ReadFile(configFile)

		if err != nil {
			return errors.New("config file load error:" + err.Error())
		}

		_, err = toml.Decode(string(configBytes), &Conf)

		if err != nil {
			return errors.New("config file parse error:" + err.Error())
		}
	}

	log.Info("config data:%v", Conf)
	return nil

}

func GetLogLvl() log.Lvl {
	switch Conf.LogLevel {
	case "DEBUG":
		return log.DEBUG
	case "INFO":
		return log.INFO
	case "WARN":
		return log.WARN
	case "ERROR":
		return log.ERROR
	case "OFF":
		return log.OFF
	}

	return log.DEBUG
}

const (
	// Template Type
	PONGO2   = "PONGO2"
	TEMPLATE = "TEMPLATE"

	// Bindata
	BINDATA = "BINDATA"

	// File
	FILE = "FILE"

	// Redis
	REDIS = "REDIS"

	// Memcached
	MEMCACHED = "MEMCACHED"

	// Cookie
	COOKIE = "COOKIE"

	// In Memory
	IN_MEMORY = "IN_MEMARY"
)