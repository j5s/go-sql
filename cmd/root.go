package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-sql/config"
	"go-sql/pkg/database"
	"go-sql/pkg/util"
	"gopkg.in/yaml.v2"
	"strings"
	"time"
)

var (
	file    string
	minRow  int
	keyword string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&file, "file", "f", "db.yaml", "config file[xx.yaml]")
	rootCmd.PersistentFlags().IntVarP(&minRow, "minRow", "", 0, "min row[xx]")
	rootCmd.PersistentFlags().StringVarP(&keyword, "keyword", "k", "name,user,pass,phone,mobile,email,card,certificate,number,addr,姓名,电话,邮箱,身份证,地址", "column keywords[xxx,xxx]")
}

var rootCmd = &cobra.Command{
	Use:               "go-sql",
	Short:             "数据库敏感信息获取工具 by zp857",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	Run: func(cmd *cobra.Command, args []string) {
		var dbConfigs config.DbConfig
		data := util.ReadFile(file)
		err := yaml.Unmarshal(data, &dbConfigs)
		if err != nil {
			fmt.Println(err)
			return
		}
		if keyword != "" {
			config.Conf.Keywords = strings.Split(keyword, ",")
		}
		fmt.Printf("[匹配关键词] %v\n\n", config.Conf.Keywords)
		for _, dbConfig := range dbConfigs.DB {
			if config.SupportDb[dbConfig.DbType] {
				runner, err := database.NewFuncMap[dbConfig.DbType](dbConfig, minRow)
				if err != nil {
					fmt.Printf("[-] %v数据库连接失败, %v:%v\n", dbConfig.DbType, dbConfig.Conn.Host, dbConfig.Conn.Port)
					continue
				}
				fmt.Printf("[+] %v数据库连接成功, %v:%v\n", dbConfig.DbType, dbConfig.Conn.Host, dbConfig.Conn.Port)
				if dbConfig.Sql == "" {
					database.RunFuncMap[dbConfig.DbType](runner)
				} else {
					runner.RunSql(dbConfig.Sql)
				}
			}
		}
	},
}

func Execute() {
	start := time.Now()
	cobra.CheckErr(rootCmd.Execute())
	fmt.Printf("spent time: %v\n", time.Since(start))
}
