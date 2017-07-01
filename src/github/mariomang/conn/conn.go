package conn

import "database/sql"
import "fmt"
import "time"
import _ "github.com/bmizerany/pq"

type PgSql struct {
	db     *sql.DB
	result *sql.Result
	rows   *sql.Rows
}

func GetPgSql() *PgSql {
	var pg *PgSql
	if pg == nil {
		globalDB, err := sql.Open("postgres", "user=postgres password=none853275286 dbname=dragonli host=123.206.44.24 sslmode=disable")
		if err != nil {
			fmt.Println(err.Error())
		}

		globalDB.SetMaxOpenConns(2000)
		globalDB.SetMaxIdleConns(1000)
		globalDB.SetConnMaxLifetime(time.Duration(time.Second * 64))

		pg = &PgSql{
			db: globalDB,
		}
	}
	return pg
}

func (pg *PgSql) Insert(sql string, params ...interface{}) sql.Result {
	stmt, err := pg.db.Prepare(sql)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	res, err := stmt.Exec(params...)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return res
}

func (pg *PgSql) Update(sql string, params ...interface{}) sql.Result {
	stmt, err := pg.db.Prepare(sql)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	res, err := stmt.Exec(params...)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return res
}

func (pg *PgSql) Query(sql string, params ...interface{}) *sql.Rows {
	db := pg.db
	rows, err := db.Query(sql, params...)
	fmt.Println(rows)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	pg.rows = rows
	return rows
}

func (pg *PgSql) Close() {
	if pg.db != nil {
		pg.db.Close()
	}
}
