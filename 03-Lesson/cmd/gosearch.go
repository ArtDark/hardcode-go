package main

import (
	"GoSearch/03-Lesson/pkg/crawler"
	"GoSearch/03-Lesson/pkg/crawler/spider"
	"GoSearch/03-Lesson/pkg/index"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//Инициализация аргументов
	sFlag := flag.String("s", "", `поиск слова на страницах "go.dev" и "golang.org"`)
	dFlag := flag.Int("d", 0, "глубина перехода по ссылкам")
	flag.Usage = help
	flag.Parse()

	//Инициализация робота
	spr := spider.New()
	//Инициализация БД индексов
	inx := index.New()

	//Список адресов
	adr := []string{"https://go.dev", "https://golang.org"}

	//Результат сканированя робота
	var data []crawler.Document

	//Добавление результатов сканирования
	for _, v := range adr {
		res, err := spr.Scan(v, *dFlag)
		if err != nil {
			log.Printf("Ошибка! %v", err)
			os.Exit(1)
		}
		data = append(data, res...)
	}
	// добавление идентификаторв
	data = spider.EnumId(data)

	// добавление индексов в БД
	if err := inx.Add(data...); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result := inx.Find(*sFlag)
	for _, v := range result {
		d := binarySearch(data, v)
		fmt.Println(d.URL)
	}

}

//help справка по использованию утилиты
func help() {
	_, _ = fmt.Printf("Использование:\n")
	_, _ = fmt.Printf("  %s [аргументы]\n", filepath.Base(os.Args[0]))
	_, _ = fmt.Printf("\n")

	_, _ = fmt.Printf("Aргументы:\n")
	flag.PrintDefaults()
}

// binarySearch поиск документа
func binarySearch(data []crawler.Document, value int) *crawler.Document {
	low := 0
	high := len(data) - 1
	for low <= high {
		mid := (low + high) / 2
		if data[mid].ID > value {
			high = mid - 1
		} else if data[mid].ID < value {
			low = mid + 1
		} else {
			return &data[mid]
		}
	}
	return nil
}
