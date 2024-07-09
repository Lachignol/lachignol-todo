package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Init() (string, error) {
	userdire, _ := os.UserHomeDir()
	dirpath := fmt.Sprintf("%s/lachignol-todo/", userdire)
	dbPath := fmt.Sprintf("%s/lachignol-todo/todos.json", userdire)

	// Vérifier si le fichier existe
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		log.Printf("Le fichier de la base de données %s n'existe pas.", dbPath)
		log.Printf("Nous procédons donc a l'installation ...")
		//si il existe pas on crée le repertoire
		err := os.Mkdir(dirpath, 0700)
		if err != nil {
			fmt.Println(err)
		}
		log.Printf("Création du repertoire %s .", dirpath)
		log.Printf("Installation terminée.")
	}
	return dbPath, nil
}

type Storage[T any] struct {
	JsonFilepath string
}

func NewStorage[t any](jsonFilepath string) *Storage[t] {
	return &Storage[t]{JsonFilepath: jsonFilepath}
}

func (s *Storage[T]) Save(data T) error {
	filedata, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err

	}
	return os.WriteFile(s.JsonFilepath, filedata, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	filedata, err := os.ReadFile(s.JsonFilepath)
	if err != nil {
		return err
	}
	return json.Unmarshal(filedata, data)

}


