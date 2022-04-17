// Package index создает структуру индексов для найденных ссылок
package index

import (
	"GoSearch/03-Lesson/pkg/crawler"
	"fmt"
	"strings"
)

// index слово из описания ссылок, значение – номера документов.
type index map[string][]int

// Db структура для храннения инвертированных индексов
type Db struct {
	index
}

func New() *Db {
	return &Db{index: index{}}
}

// Add добавляет в структуру Db слова из описания ссылок и идентификатор
func (d *Db) Add(doc ...crawler.Document) error {
	if d == nil {
		return fmt.Errorf("index database is %v", d)
	}
	for _, v := range doc {
		words := strings.Split(v.Title, " ")
		for _, w := range words {
			s := strings.ToLower(w)
			if d.exist(s, v.ID) {
				continue
			}
			d.index[s] = append(d.index[s], v.ID)
		}
	}
	return nil
}

// exist проверяет наличие слова в базе индексов
func (d *Db) exist(s string, id int) bool {
	itm := d.index[s]

	return contains(itm, id)
}

// Find поиск слова в Бд индексов. Возвращает идет идентификаторы документов.
func (d *Db) Find(s string) []int {
	lowS := strings.ToLower(s)
	return d.index[lowS]
}

// contains сообщает имеется ли элемент в списке
func contains(s []int, item int) bool {
	for _, v := range s {
		if v == item {
			return true
		}
	}

	return false
}
