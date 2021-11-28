package database

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"go-sql/pkg/util"
)

type Runner struct {
	MinRow 	int
	Engine *xorm.Engine
}

func NewRunner(driver, dsn string, minRow int) (*Runner, error) {
	engine, err := xorm.NewEngine(driver, dsn)
	if err != nil {
		return nil, err
	}
	return &Runner{
		Engine: engine,
		MinRow: minRow,
	}, nil
}

func (r *Runner) RunSql(sql string) {
	fmt.Printf("[SQL语句] %v\n", sql)
	rows, err := r.Engine.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	if len(rows) > 0 {
		// 获取字段名
		var columns []string
		for k := range rows[0] {
			columns = append(columns, k)
		}
		res := "[COLUMN] "
		for _, column := range columns {
			res += column + " | "
		}
		res = res[:len(res)-3]
		fmt.Println(res)
		for _, result := range rows {
			res = "[ROW] "
			for _, column := range columns {
				res += util.Strval(result[column]) + " "
			}
			res = res[:len(res)-1]
			fmt.Println(res)
		}
	} else {
		fmt.Println("[WARN] 结果为空")
	}
}

