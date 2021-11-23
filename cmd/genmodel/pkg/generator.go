package pkg

import (
	"bytes"
	"database/sql"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type TableColumn struct {
	OrdinalPosition uint16 `db:"ORDINAL_POSITION"` // position
	ColumnName      string `db:"COLUMN_NAME"`      // name
	ColumnType      string `db:"COLUMN_TYPE"`      // column_type
	DataType        string `db:"DATA_TYPE"`        // data_type
	FieldName       string
	FieldType       string
	ColumnKey       sql.NullString `db:"COLUMN_KEY"`     // key
	IsNullable      string         `db:"IS_NULLABLE"`    // nullable
	Extra           sql.NullString `db:"EXTRA"`          // extra
	ColumnComment   string         `db:"COLUMN_COMMENT"` // comment
	ColumnDefault   sql.NullString `db:"COLUMN_DEFAULT"` // default value
}

type Values struct {
	PkgName    string
	StructName string
	Fields     []TableColumn
}

func GeneratorModel(tableName string, structName string, columns []TableColumn) {
	values := Values{tableName, structName, columns}
	buf := new(bytes.Buffer)
	err := outputTemplate.Execute(buf, values)
	if err != nil {
		panic(err)
	}

	//格式化代码
	formattedOutput, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	buf = bytes.NewBuffer(formattedOutput)

	outDir := "internal/model/" + strings.ToLower(structName)
	err = os.Mkdir(outDir, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	//输出文件
	filename := fmt.Sprintf("%s/%s.go", outDir, strings.ToLower(structName))
	if err := ioutil.WriteFile(filename, buf.Bytes(), 0777); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("  └── file : ", fmt.Sprintf("%s", filename))
}
