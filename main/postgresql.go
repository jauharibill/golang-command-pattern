package main

type postgresqlCommand struct {
	database database
}

func (p *postgresqlCommand) connect() {
	p.database.postgresql()
}
