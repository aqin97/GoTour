package cmd

import (
	"GoTour/commandlinetool/internal/sql2struct"
	"log"

	"github.com/spf13/cobra"
)

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql转换和处理",
	Long:  "sql转换和处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql转换至Go结构体",
	Long:  "sql转换至Go结构体",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			Username: username,
			Password: password,
			Charset:  charset,
		}
		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err: %v", err)
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", err)
		}
		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssembleColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "u", "", "请输入数据库用户名")
	sql2structCmd.Flags().StringVarP(&password, "password", "p", "", "请输入数据库密码")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "47.100.21.38", "请输入数据库的HOST")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "c", "utf8mb4", "请输入数据库字符编码类型")
	sql2structCmd.Flags().StringVarP(&dbType, "dbtype", "t", "mysql", "请输入数据库类型")
	sql2structCmd.Flags().StringVarP(&dbName, "dbname", "n", "", "请输入数据库名称")
	sql2structCmd.Flags().StringVarP(&tableName, "tablename", "T", "", "请输入数据表名称")
}
