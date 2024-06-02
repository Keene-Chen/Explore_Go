package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 数据库连接字符串
	// 格式: username:password@protocol(address)/dbname?param=value
	dsn := "root:mysql_WYGjtt@tcp(gcloud.keenechen.cn:3306)/mysql"

	// 打开数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 检查数据库连接是否成功
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database")

	// 执行查询
	query := "SELECT * FROM user"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 获取列名称
	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	// 创建表格
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(columns)

	// 预分配一个切片来存放行数据
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// 遍历查询结果
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal(err)
		}

		// 将每行的数据转换为字符串切片
		row := make([]string, len(columns))
		for i, col := range values {
			if col == nil {
				row[i] = "NULL"
			} else {
				row[i] = string(col)
			}
		}
		table.Append(row)
	}

	// 检查是否有任何行扫描错误
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// 渲染表格
	table.Render()
}