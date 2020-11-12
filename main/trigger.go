package main

type trigger struct {
	command command
}

func (t *trigger) execute() {
	t.command.connect()
}
