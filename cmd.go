package main

import (
	"flag"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewcmdFlags() *CmdFlags {
	cf := CmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "Ajouter une tache")
	flag.StringVar(&cf.Edit, "edit", "", "Editer une tache existante spécifié par son index")
	flag.IntVar(&cf.Del, "del", -1, "Supprimer une tache spécifié par son index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "compléter/décompléter une tache existante spécifié par son index")
	flag.BoolVar(&cf.List, "list", false, "compléter/décompléter une tache existante spécifié par son index")

	flag.Parse()

	return &cf

}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.Addtodo(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			println("Error invalid format for edit.Please use id:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			println("Invalid index")
			os.Exit(1)
		}
		todos.EditTodo(index, parts[1])
	case cf.Toggle != -1:
		todos.Toggle(cf.Toggle)
	case cf.Del != -1:
		todos.DeleteTodo(cf.Del)

	default:
		println("Usage of todo-app:\n\n-add <'Titre'>\nAjouter une tache.\n\n-del <Id>\nSupprimer une tache spécifié par son index.\n\n-edit <Id:'Nouveau titre'>\nEditer une tache existante spécifié par son index puis entrer son nouveau titre.\n\n-toggle <Id>\nCompléter/décompléter une tache existante spécifié par son index.\n\n-list\nlister les taches.")
	}

}
