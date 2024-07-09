package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	Filename string
}

func NewStorage[t any](filename string) *Storage[t] {
	return &Storage[t]{Filename: filename}
}

func (s *Storage[T]) Save(data T) error {
	filedata, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err

	}
	return os.WriteFile(s.Filename, filedata, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	filedata, err := os.ReadFile(s.Filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(filedata, data)

}
