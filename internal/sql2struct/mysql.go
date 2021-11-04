package sql2struct

import (
	"database/sql"
	"errors"
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

//获取表的信息
func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	query := "SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?"
	rows, err := m.DBEngine.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("没有数据")
	}
	defer rows.Close()

	var columns []*TableColumn
	for rows.Next() {
		var colunm TableColumn
		err := rows.Scan(&colunm.ColumnName, &colunm.DataType, &colunm.ColumnKey, &colunm.IsNullable, &colunm.ColumnType, &colunm.ColumnComment)
		if err != nil {
			return nil, err
		}

		columns = append(columns, &colunm)
	}
	return columns, nil
}

//表字段类型映射
var DBTypeToStructType = map[string]string{
	"int":        "int32",
	"tinyint":    "int8",
	"smallint":   "int",
	"mediumint":  "int64",
	"bigint":     "int64",
	"bit":        "int",
	"bool":       "bool",
	"enum":       "string",
	"set":        "string",
	"varchar":    "string",
	"char":       "string",
	"tinytext":   "string",
	"mediumtext": "string",
	"text":       "string",
	"longtext":   "string",
	"blob":       "string",
	"tinyblob":   "string",
	"mediumblob": "string",
	"longblob":   "string",
	"date":       "time.Time",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"time":       "time.Time",
	"float":      "float64",
	"double":     "float64",
}
