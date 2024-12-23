package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	Filename string
}

func Newstorage[T any](fname string) *Storage[T] {
	return &Storage[T]{Filename: fname}
}

func (s *Storage[T]) Save(data T) error {
	fdata, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		return err
	}

	return os.WriteFile(s.Filename, fdata, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	fdata, err := os.ReadFile(s.Filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(fdata, data)
}