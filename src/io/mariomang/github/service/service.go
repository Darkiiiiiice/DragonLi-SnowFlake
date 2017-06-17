package service

import (
	"fmt"
	"io/mariomang/github/dao"
	"io/mariomang/github/domain"
	"time"
)

func GenrateIDService(request *domain.RequestDomain) string {

	pg := dao.GetPgSql()
	rows := pg.Query("SELECT * FROM dragonli")

	fmt.Println(rows)
	if rows == nil {
		return ""
	}

	for rows.Next() {
		var busid string
		var maxid int64
		var step int64
		var desc string
		var updatedate time.Duration
		rows.Scan(&busid, &maxid, &step, &desc, &updatedate)
		columns, _ := rows.Columns()
		fmt.Println(rows.Columns())
		fmt.Printf("%s=%s,%s=%d,%s=%d,%s=%s,%s=%v\n", columns[0], busid, columns[1], maxid, columns[2], step, columns[3], desc, columns[4], updatedate)
	}

	pg.Close()

	return "HelloWorld"
}
