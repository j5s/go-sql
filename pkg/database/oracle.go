package database

import (
	"fmt"
	_ "github.com/mattn/go-oci8"
	"go-sql/config"
	"go-sql/pkg/util"
	"strconv"
	"strings"
)

func NewOracle(db *config.DB, minRow int) (*Runner, error) {
	dsn := fmt.Sprintf("%v/%v@%v:%v/%v", db.Conn.User, db.Conn.Pass, db.Conn.Host, db.Conn.Port, db.Conn.DbName)
	return NewRunner("oci8", dsn, minRow)
}

func RunOracleDefault(r *Runner) {
	var sql string
	// oracle没有库名，查询表空间
	var dbNames []string
	sql = config.Conf.Oracle.QueryDb
	rows, err := r.Engine.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	for _, result := range rows {
		dbNames = append(dbNames, util.Strval(result["NAME"]))
	}
	fmt.Printf("[数据库] %v\n", dbNames)
	// 获取表行数并统计
	var genLike string
	for _, keyword := range config.Conf.Keywords {
		genLike += strings.Replace(config.Conf.Oracle.GenLike, "{keyword}", keyword, -1)
	}
	sql = strings.Replace(config.Conf.Oracle.QueryColumn, "{genLike}", genLike[:len(genLike)-1], -1)
	fmt.Printf("[所有表匹配关键词] %v\n", sql)
	rows, err = r.Engine.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	//字段名
	fmt.Println("[COLUMN] 所有者 | 表名 | 字段名 | 行数")
	for _, result := range rows {
		row, _ := strconv.Atoi(util.Strval(result["NUM_ROWS"]))
		if row < r.MinRow {
			continue
		}
		fmt.Printf("[ROW] %v | %v | %v | %v\n", util.Strval(result["OWNER"]), util.Strval(result["TABLE_NAME"]), util.Strval(result["COLUMN_NAME"]), util.Strval(result["NUM_ROWS"]))
	}
	fmt.Println()
	// 数据库连接统计
	sql = config.Conf.Oracle.QueryConnect
	fmt.Printf("[数据库连接统计] %v\n", sql)
	rows, err = r.Engine.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("[COLUMN] 用户名 | 系统用户名 | 机器名")
	for _, result := range rows {
		fmt.Printf("[ROW] %v | %v | %v\n", util.Strval(result["USERNAME"]), util.Strval(result["OSUSER"]), util.Strval(result["MACHINE"]))
	}
	fmt.Println()
}
