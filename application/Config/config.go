package Config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Server
	DataBase
}
type Server struct {
	Port int `yaml:"port"`
}
type DataBase struct {
	Host  string `yaml:"host"`
	Table string `yaml:"table"`
	User  string `yaml:"user"`
	Pass  string `yaml:"pass"`
	Port  string `yaml:"port"`
}

func (c Config) GetConf() Config {
	conf := Config{}
	file, err := ioutil.ReadFile("application/Config/config.yml")
	if err != nil {
		panic("读取配置文件失败")
	}

	if err := yaml.Unmarshal(file, &conf); err != nil {
		panic("格式化数据失败")
	}
	return conf
}
