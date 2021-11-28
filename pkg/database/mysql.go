package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-sql/config"
	"go-sql/pkg/util"
	"strconv"
	"strings"
)

func NewMysql(db *config.DB, minRow int) (*Runner, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", db.Conn.User, db.Conn.Pass, db.Conn.Host, db.Conn.Port, db.Conn.DbName)
	return NewRunner("mysql", dsn, minRow)
}

func RunMysqlDefault(r *Runner) {
	var sql string
	// 获取全部数据库
	var dbNames []string
	sql = config.Conf.Mysql.QueryDb
	rows, err := r.Engine.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	for _, result := range rows {
		db := util.Strval(result["DB"])
		if util.IsContain(config.Conf.Mysql.DefaultDb, db) {
			continue
		}
		dbNames = append(dbNames, db)
	}
	fmt.Printf("[数据库] %v\n", dbNames)
	// 获取表行数并统计
	var genLike string
	for _, keyword := range config.Conf.Keywords {
		genLike += strings.Replace(config.Conf.Mysql.GenLike, "{keyword}", keyword, -1)
	}
	sql = strings.Replace(config.Conf.Mysql.QueryColumn, "{genLike}", genLike[:len(genLike)-1], -1)
	fmt.Printf("[所有表匹配关键词] %v\n", sql)
	rows, err = r.Engine.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	//字段名
	fmt.Println("[COLUMN] 库名 | 表名 | 字段名 | 行数")
	for _, result := range rows {
		db := util.Strval(result["TABLE_SCHEMA"])
		if util.IsContain(config.Conf.Mysql.DefaultDb, db) {
			continue
		}
		row, _ := strconv.Atoi(util.Strval(result["TABLE_ROWS"]))
		if row < r.MinRow {
			continue
		}
		fmt.Printf("[ROW] %v | %v | %v | %v\n", db, util.Strval(result["TABLE_NAME"]), util.Strval(result["COLUMN_NAME"]), row)
	}
	fmt.Println()
	// 数据库连接统计
	sql = config.Conf.Mysql.QueryConnect
	fmt.Printf("[数据库连接统计] %v\n", sql)
	rows, err = r.Engine.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("[COLUMN] Id | 用户 | 主机 | 数据库")
	for _, result := range rows {
		fmt.Printf("[ROW] %v | %v | %v | %v\n", util.Strval(result["Id"]), util.Strval(result["User"]), util.Strval(result["Host"]), util.Strval(result["db"]))
	}
	fmt.Println()
}
