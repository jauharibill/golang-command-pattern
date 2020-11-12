package main

import "fmt"

type dbms struct {
	title string `json:"title"`
}

func (d *dbms) mysql() {
	d.title = "this is sql script to connect with mysql database"
	fmt.Println(d.title)
}

func (d *dbms) postgresql() {
	d.title = "postgresql"
	fmt.Println(d.title)
}

func (d *dbms) sqlserver() {
	d.title = "sqlserver"
	fmt.Println(d.title)
}
