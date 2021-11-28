package config

import (
	"embed"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

//go:embed config.yaml
var f embed.FS

var Conf = new(InitConfig)

type InitConfig struct {
	Keywords    []string  `yaml:"keywords"`
	Mssql *SqlConfig `yaml:"mssql"`
	Mysql SqlConfig `yaml:"mysql"`
	Oracle SqlConfig `yaml:"oracle""`
	Postgres SqlConfig `yaml:"postgres"`
}

type SqlConfig struct {
	QueryDb      string   `yaml:"query_db"`
	DefaultDb    []string `yaml:"default_db"`
	GenLike      string   `yaml:"gen_like"`
	QueryColumn  string   `yaml:"query_column"`
	QueryConnect string   `yaml:"query_connect"`
}

func init() {
	data, _ := f.ReadFile("config.yaml")
	err := yaml.Unmarshal(data, &Conf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
