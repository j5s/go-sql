package config

type DbConfig struct {
	DB []*DB `yaml:"db"`
}

type DB struct {
	DbType string `yaml:"db_type"`
	Conn   struct {
		Host   string `yaml:"host"`
		Port   int    `yaml:"port"`
		DbName string `yaml:"db_name"`
		User   string `yaml:"user"`
		Pass   string `yaml:"pass"`
	} `yaml:"conn"`
	Sql string `yaml:"sql"`
}

var SupportDb = map[string]bool{
	"mysql":    true,
	"mssql":    true,
	"oracle":   true,
	"postgres": true,
}