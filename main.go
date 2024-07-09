package main

var dbPath string

func init() {
	dbPath, _ = Init()
}

func main() {
	todos := Todos{}
	storage := NewStorage[Todos](dbPath)
	storage.Load(&todos)
	cmdFlags := NewcmdFlags()
	cmdFlags.Execute(&todos)
	storage.Save(todos)

}
