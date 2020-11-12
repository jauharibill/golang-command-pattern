package main

type database interface {
	mysql()
	postgresql()
	sqlserver()
}
