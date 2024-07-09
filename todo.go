package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title        string
	Completed    bool
	Created_at   time.Time
	Completed_at *time.Time
}

type Todos []Todo

func (todos *Todos) Addtodo(title string) {
	todo := Todo{
		Title:        title,
		Completed:    false,
		Created_at:   time.Now(),
		Completed_at: nil,
	}

	*todos = append(*todos, todo)

}

func (todos *Todos) CheckTodo(index int) error {
	if index < 0 || index >= len(*todos) {
		error := errors.New("invalid index")
		fmt.Println(error)
		return error
	}
	return nil
}

func (Todos *Todos) DeleteTodo(index int) error {
	t := *Todos
	if err := t.CheckTodo(index); err != nil {
		return err
	}
	*Todos = append(t[:index], t[index+1:]...)
	return nil
}

func (Todos *Todos) Toggle(index int) error {
	t := *Todos
	if err := t.CheckTodo(index); err != nil {

		return err
	}
	iscompleted := t[index].Completed

	if !iscompleted {
		Completed_at := time.Now()
		t[index].Completed_at = &Completed_at
	} else {
		t[index].Completed_at = nil
	}
	t[index].Completed = !iscompleted

	return nil
}

func (Todos *Todos) EditTodo(index int, title string) error {
	t := *Todos
	if err := t.CheckTodo(index); err != nil {

		return err
	}

	t[index].Title = title

	return nil
}

func (Todos *Todos) print() {
	t := *Todos
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "titre", "completed", "created_at", "completed_at")
	for idx, todo := range t {
		completed := "❌"
		completedAt := ""
		if todo.Completed {
			completed = "✅"
		}
		if todo.Completed_at != nil {
			completedAt = todo.Completed_at.Format(time.RFC850)
		}
		table.AddRow(strconv.Itoa(idx), todo.Title, completed, todo.Created_at.Format(time.RFC850), completedAt)

	}
	table.Render()
}

func (todos *Todos) Uninstall() {
	userdire, _ := os.UserHomeDir()
	dirpath := fmt.Sprintf("%s/lachignol-todo/", userdire)
	dbPath := fmt.Sprintf("%s/lachignol-todo/todos.json", userdire)

	err := os.Remove(dbPath)
	if err != nil {
		log.Fatal(err)

	}

	fmt.Printf("Supression du fichier %v \n\n", dbPath)

	e := os.Remove(dirpath)
	if e != nil {
		log.Fatal(e)

	}
	fmt.Printf("Supression du repertoire %v\n\nIl ne vous reste plus qu'a supprimer le binaire 'lachignol-todo' dans le repertoire $HOME/go/bin/ .\n", dirpath)

}
