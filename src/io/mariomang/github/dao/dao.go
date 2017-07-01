package dao

import (
	"fmt"
	"github/mariomang/conn"
	"time"
)

func GetMaxIDAndStepByWorkID(workID int64) (int64, int64) {
	var maxID, step int64
	pg := conn.GetPgSql()

	defer pg.Close()

	sql := "SELECT max_id maxID,step step FROM dragonli where work_id=$1"
	rows := pg.Query(sql, workID)
	fmt.Println("execute sql:" + sql)

	if rows.Next() {
		rows.Scan(&maxID, &step)
	}
	return maxID, step
}

func HasWorkID(workID int64) bool {
	pg := conn.GetPgSql()

	defer pg.Close()

	sql := "SELECT work_id FROM dragonli where work_id=$1"
	rows := pg.Query(sql, workID)
	fmt.Println("execute sql:" + sql)
	if rows.Next() {
		return true
	}
	return false
}

func InsertByWorkID(workID int64, desc string) {
	pg := conn.GetPgSql()

	defer pg.Close()

	sql := "INSERT INTO dragonli (work_id, descp, update_time)VALUES ($1,$2,$3)"
	// fmt.Println(t)
	result := pg.Insert(sql, workID, desc, time.Now())
	// fmt.Println("execute sql:" + sql)

	_, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
	}
}
