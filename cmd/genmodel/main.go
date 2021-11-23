package main

import (
	"eicesoft/web-demo/config"
	"eicesoft/web-demo/pkg/db"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"os"
	"strings"

	"eicesoft/web-demo/cmd/genmodel/pkg"
)

var (
	table      string
	structName string
)

func init() {
	flagTable := flag.String("table", "", "[Required] The name of the db table name\n")
	flagstructName := flag.String("struct", "", "[Required] The name of the db table model\n")

	if !flag.Parsed() {
		flag.Parse()
	}

	if *flagTable == "" {
		flag.Usage()
		os.Exit(1)
	}

	table = *flagTable
	structName = *flagstructName
}

// capitalize 格式化字符串
func capitalize(s string) string {
	var upperStr string
	chars := strings.Split(s, "_")
	for _, val := range chars {
		vv := []rune(val)
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				if vv[i] >= 97 && vv[i] <= 122 {
					vv[i] -= 32
					upperStr += string(vv[i])
				}
			} else {
				upperStr += string(vv[i])
			}
		}
	}
	return upperStr
}

// textType 获得数据库对应go中类型
func textType(s string) string {
	var mysqlTypeToGoType = map[string]string{
		"tinyint":    "int32",
		"smallint":   "int32",
		"mediumint":  "int32",
		"int":        "int32",
		"integer":    "int64",
		"bigint":     "int64",
		"float":      "float64",
		"double":     "float64",
		"decimal":    "float64",
		"date":       "string",
		"time":       "string",
		"year":       "string",
		"datetime":   "time.Time",
		"timestamp":  "time.Time",
		"char":       "string",
		"varchar":    "string",
		"tinyblob":   "string",
		"tinytext":   "string",
		"blob":       "string",
		"text":       "string",
		"json":       "string",
		"mediumblob": "string",
		"mediumtext": "string",
		"longblob":   "string",
		"longtext":   "string",
	}
	return mysqlTypeToGoType[s]
}

func main() {
	dbRepo, err := db.New()
	if err != nil {
		fmt.Printf("数据库连接失败", zap.Error(err))
	}
	db := dbRepo.GetDbR()
	sqlTableColumn := fmt.Sprintf("SELECT `ORDINAL_POSITION`,`COLUMN_NAME`,`COLUMN_TYPE`,`DATA_TYPE`,`COLUMN_KEY`,"+
		"`IS_NULLABLE`,`EXTRA`,`COLUMN_COMMENT`,`COLUMN_DEFAULT` FROM `information_schema`.`columns` WHERE `table_schema`= "+
		"'%s' AND `table_name`= '%s' ORDER BY `ORDINAL_POSITION` ASC",
		config.Get().MySQL.Read.Name, table)
	rows, err := db.Raw(sqlTableColumn).Rows()

	if err != nil {
		fmt.Printf("execute query table column action error, detail is [%v]\n", err.Error())
	}

	i := 0
	columns := make([]pkg.TableColumn, 0)
	for rows.Next() {
		i++
		var column pkg.TableColumn
		err = rows.Scan(
			&column.OrdinalPosition,
			&column.ColumnName,
			&column.ColumnType,
			&column.DataType,
			&column.ColumnKey,
			&column.IsNullable,
			&column.Extra,
			&column.ColumnComment,
			&column.ColumnDefault)
		if err != nil {
			fmt.Printf("query table column scan error, detail is [%v]\n", err.Error())
		}
		column.FieldName = capitalize(column.ColumnName)
		column.FieldType = textType(column.DataType)
		columns = append(columns, column)
		//fmt.Printf("%s\n", column.ColumnName)
	}
	pkg.GeneratorModel(table, structName, columns)

	defer rows.Close()
}
