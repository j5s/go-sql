package database

import (
	"go-sql/config"
)

type NewFunc func(db *config.DB, minRow int) (*Runner, error)

type RunFunc func(runner *Runner) ()

var (
	NewFuncMap map[string]NewFunc
	RunFuncMap map[string]RunFunc
)

func init() {
	NewFuncMap = make(map[string]NewFunc)
	NewFuncMap["mssql"] = NewMssql
	NewFuncMap["mysql"] = NewMysql
	NewFuncMap["oracle"] = NewOracle
	NewFuncMap["postgres"] = NewPostgres

	RunFuncMap = make(map[string]RunFunc)
	RunFuncMap["mssql"] = RunMssqlDefault
	RunFuncMap["mysql"] = RunMysqlDefault
	RunFuncMap["oracle"] = RunOracleDefault
	RunFuncMap["postgres"] = RunPostgresDefault
}