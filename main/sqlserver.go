package main

type sqlserverCommand struct {
	database database
}

func (s *sqlserverCommand) connect() {
	s.database.sqlserver()
}
