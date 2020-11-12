package main

func main() {

	comm := &mysqlCommand{database: &dbms{}}
	triger := trigger{command: comm}
	triger.execute()
}
