package sql2struct

import (
	"GoTour/commandlinetool/internal/word"
	"fmt"
	"os"
	"text/template"
)

const structTpl = `type {{ .TableName | ToCamelCase }} struct {
{{ range .Columns }} {{ $length := len .Comment }} {{ if gt $length 0 }}//{{.Comment }} {{ else }}// {{ .Name }} {{ end }}
{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }} {{ .Name | ToCamelCase }}	{{ .Type }}	{{.Tag}} {{ else }} {{.Name}}{{ end }}{{ end }}
}

func (model {{ .TableName | ToCamelCase }}) TableName() string {
	return "{{ .TableName }}"
}`

type StructTemplate struct {
	structTpl string
}

//存字段信息
type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

//最终用于渲染上述模板的信息，包括数据表名，以及所有表->结构体的字段信息
type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTpl}
}

//处理数据库类型到go结构体字段以及tag
func (t *StructTemplate) AssembleColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tag := fmt.Sprintf(" "+"json:"+"\"%s\""+" ", column.ColumnName)
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     tag,
			Comment: column.ColumnComment,
		})
	}

	return tplColumns
}

//模板中有一些自定义函数，下面的函数负责处理模板自定义函数和模板对象
func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{"ToCamelCase": word.UnderscoreToUpperCamelCase}).Parse(t.structTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}
	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}

	return nil
}
