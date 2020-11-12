package main

type mysqlCommand struct {
	database database
}

func (d *mysqlCommand) connect() {
	d.database.mysql()
}
