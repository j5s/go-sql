package database

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"go-sql/config"
	"go-sql/pkg/util"
	"strconv"
	"strings"
)

func NewMssql(db *config.DB, minRow int) (*Runner, error) {
	dsn := fmt.Sprintf("sqlserver://%v:%v@%v:%v?database=%v&connection+timeout=1", db.Conn.User, db.Conn.Pass, db.Conn.Host, db.Conn.Port, db.Conn.DbName)
	return NewRunner("mssql", dsn, minRow)
}

func RunMssqlDefault(r *Runner) {
	var sql string
	// 获取全部数据库
	var dbNames []string
	sql = config.Conf.Mssql.QueryDb
	rows, err := r.Engine.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	for _, result := range rows {
		db := util.Strval(result["DB"])
		if util.IsContain(config.Conf.Mssql.DefaultDb, db) {
			continue
		}
		dbNames = append(dbNames, db)
	}
	fmt.Printf("[数据库] %v\n", dbNames)
	// 获取表行数并统计
	var genLike string
	for _, keyword := range config.Conf.Keywords {
		genLike += strings.Replace(config.Conf.Mssql.GenLike, "{keyword}", keyword, -1)
	}
	sql = strings.Replace(config.Conf.Mssql.QueryColumn, "{genLike}", genLike[:len(genLike)-4], -1)
	fmt.Println("[COLUMN] 库名 | 表名 | 字段名 | 行数")
	for _, dbName := range dbNames {
		query := strings.Replace(sql, "{db}", dbName, -1)
		fmt.Printf("[所有表匹配关键词] %v\n", query)
		rows, err = r.Engine.Query(query)
		if err != nil {
			fmt.Println(err)
		}
		for _, result := range rows {
			row, _ := strconv.Atoi(util.Strval(result["TABLE_ROWS"]))
			if row < r.MinRow {
				continue
			}
			fmt.Printf("[ROW] %v | %v | %v | %v\n", dbName, util.Strval(result["TABLE_NAME"]), util.Strval(result["COLUMN_NAME"]), row)
		}
	}
	fmt.Println()
	// 数据库连接统计
	sql = config.Conf.Mssql.QueryConnect
	fmt.Printf("[数据库连接统计] %v\n", sql)
	fmt.Println("[COLUMN] 客户端IP | 本地IP")
	rows, err = r.Engine.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	for _, result := range rows {
		fmt.Printf("[ROW] %v | %v\n", util.Strval(result["client_net_address"]), util.Strval(result["local_net_address"]))
	}
	fmt.Println()
}
