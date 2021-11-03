package sql2struct

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//数据库连接的核心
type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

//所连接的数据库信息
type DBInfo struct {
	DBType   string
	Host     string
	Username string
	Password string
	Charset  string
}

//用于存放mysql中information_schema数据库中的COLUMNS数据表
type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}

//数据库连接函数
func (m *DBModel) Connect() error {
	var err error
	s := "%s:%s@tcp(%s)/information_schema?charset=%s&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		s,
		m.DBInfo.Username,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset,
	)
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}
	return nil
}
