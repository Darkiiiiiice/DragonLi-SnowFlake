package dao

import (
	"io/mariomang/github/utils"
)

func GetMaxIDAndStepByWorkID(workID string) (int64, int64) {
	var maxID, step int64
	pg := utils.GetPgSql()
	sql := "SELECT max_id maxID,step step FROM dragonli where work_id=$1"
	rows := pg.Query(sql, workID)
	if rows.Next() {
		rows.Scan(&maxID, &step)
	}
	pg.Close()
	return maxID, step
}
