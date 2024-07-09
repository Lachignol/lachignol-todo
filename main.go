package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	cmdFlags:=NewcmdFlags()
	cmdFlags.Execute(&todos)
	storage.Save(todos)

}
