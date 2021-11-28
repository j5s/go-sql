package database

import (
	"fmt"
	_ "github.com/lib/pq"
	"go-sql/config"
	"go-sql/pkg/util"
	"strings"
)

func NewPostgres(db *config.DB, minRow int) (*Runner, error) {
	// 初始化连接
	dsn :=  fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v&connect_timeout=1", db.Conn.User, db.Conn.Pass, db.Conn.Host, db.Conn.Port, db.Conn.DbName, "disable")
	return NewRunner("postgres", dsn, minRow)
}

func RunPostgresDefault(r *Runner) {
	// 获取数据库
	var dbNames []string
	sql := config.Conf.Postgres.QueryDb
	rows, err := r.Engine.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	for _, result := range rows {
		dbNames = append(dbNames, util.Strval(result["datname"]))
	}
	fmt.Printf("[数据库] %v\n", dbNames)
	// 数据表行数统计排序
	var genLike string
	for _, keyword := range config.Conf.Keywords {
		genLike += strings.Replace(config.Conf.Postgres.GenLike, "{keyword}", keyword, -1)
	}
	sql = strings.Replace(config.Conf.Postgres.QueryColumn, "{genLike}", genLike[:len(genLike)-4], -1)
	fmt.Printf("[所有表匹配关键词] %v\n", sql)
	rows, err = r.Engine.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	//字段名
	fmt.Println("[COLUMN] 库名 | 表名 | 字段名 | 行数")
	for _, result := range rows {
		fmt.Printf("[ROW] %v | %v | %v | %v | %v\n", util.Strval(result["database_name"]), util.Strval(result["table_schema"]), util.Strval(result["table_name"]), util.Strval(result["column_name"]), util.Strval(result["rows"]))
	}
	fmt.Println()
	// 数据库连接统计
	sql = config.Conf.Postgres.QueryConnect
	fmt.Printf("[数据库连接统计] %v\n", sql)
	rows, err = r.Engine.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("[COLUMN] 库名 | 用户名 | 客户端IP")
	for _, result := range rows {
		if util.Strval(result["datname"]) == "" {
			continue
		}
		fmt.Printf("[ROW] %v | %v | %v\n", util.Strval(result["datname"]), util.Strval(result["usename"]), util.Strval(result["client_addr"]))
	}
	fmt.Println()
}
