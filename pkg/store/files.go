// Package store позволяет сохранять и получать доступ к сохраненным данным
package store

import (
	"GoSearch/pkg/crawler"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

type Service struct {
	data io.ReadWriter
}

const fileName = "data"

//Save сохраняет результы поиска в формате JSON
func (s *Service) Save(d []crawler.Document) error {
	b, err := json.Marshal(d)
	if err != nil {
		return err
	}

	_, err = s.data.Write(b)
	if err != nil {
		return err
	}
	return nil
}

// New создает файл для сохранения результата поиска.
func New(path string) (*Service, error) {
	p := filepath.Dir(path)
	var str Service

	f, err := os.OpenFile(filepath.Join(p, fileName), os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		return nil, err
	}
	str.data = f

	return &str, nil
}

// Data используется для получения предварительно сохраненных документов из файла.
func (s *Service) Data() ([]crawler.Document, error) {

	var d []crawler.Document

	if err := json.NewDecoder(s.data).Decode(&d); err != nil {
		return nil, err
	}

	return d, nil

}
