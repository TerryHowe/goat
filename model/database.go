package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

type ListIterator struct {
	rows *sql.Rows
	err  error
}

func (iter *ListIterator) scan(v ...interface{}) bool {
	if iter.rows == nil {
		return false
	}
	if !iter.rows.Next() {
		iter.rows.Close()
		iter.rows = nil
		return false
	}
	err := iter.rows.Scan(v...)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (iter *ListIterator) close() bool {
	if iter.rows == nil {
		return false
	}
	iter.rows.Close()
	iter.rows = nil
	return true
}

type Database struct {
	context *sql.DB
	err     error
}

func (d *Database) open() bool {
	database_type := "mysql"
	uri := "root:asdf1234@tcp(127.0.0.1:3306)/goat"
	d.context, d.err = sql.Open(database_type, uri)
	if d.err != nil {
		log.Fatal(d.err)
	}
	d.err = d.context.Ping()
	if d.err != nil {
		log.Fatal(d.err)
		return false
	}
	return true
}

func (d *Database) list(table string) ListIterator {
	var iterator ListIterator
	iterator.rows, d.err = d.context.Query("select * from clusters")
	if d.err != nil {
		log.Fatal(d.err)
		iterator.rows = nil
		return iterator
	}
	return iterator
}

func (d *Database) create(table string, fields []string, values []sql.NullString) int64 {
	var stmt *sql.Stmt
	var res sql.Result
	fieldstr := strings.Join(fields, ",")
	valuestr := ""
	for _, element := range values {
		if len(valuestr) > 0 {
			valuestr += ","
		}
		valuestr += "'" + element.String + "'"
	}
	cmd := fmt.Sprintf("insert into %s(%s) values(%s)", table, fieldstr, valuestr)
	fmt.Printf(cmd)
	stmt, d.err = d.context.Prepare(cmd)
	if d.err != nil {
		log.Fatal(d.err)
		return 0
	}
	res, d.err = stmt.Exec()
	if d.err != nil {
		log.Fatal(d.err)
		return 0
	}
	var lastId int64
	lastId, d.err = res.LastInsertId()
	if d.err != nil {
		log.Fatal(d.err)
		return 0
	}
	return lastId
}

func (d *Database) delete(table string, identifier int64) bool {
	var stmt *sql.Stmt
	cmd := fmt.Sprintf("delete from %s where id=%d", table, identifier)
	fmt.Printf(cmd)
	stmt, d.err = d.context.Prepare(cmd)
	if d.err != nil {
		log.Fatal(d.err)
		return false
	}
	_, d.err = stmt.Exec()
	if d.err != nil {
		log.Fatal(d.err)
		return false
	}
	return true
}

func (d *Database) close() bool {
	d.err = d.context.Close()
	if d.err != nil {
		log.Fatal(d.err)
		return false
	}
	return true
}
