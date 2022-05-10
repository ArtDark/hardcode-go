package main

import (
	"GoSearch/pkg/crawler"
	"GoSearch/pkg/crawler/spider"
	"GoSearch/pkg/index"
	"GoSearch/pkg/store"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	exPath, err := os.Executable()
	if err != nil {
		log.Fatalf("ошибка! Неудается получить путь дирректории с исполняемым файлом. %v", err)
	}

	sFlag := flag.String("s", "", `поиск слова на страницах "go.dev" и "golang.org"`)
	dFlag := flag.Int("d", 0, "глубина перехода по ссылкам")

	flag.Usage = help
	flag.Parse()

	spr := spider.New()
	inx := index.New()

	adr := []string{"https://go.dev", "https://golang.org"}

	var data []crawler.Document

	str, err := store.New(exPath)
	if err != nil {
		log.Println(err)
	}

	data, err = str.Data()
	if err != nil {
		log.Println(err)
	}

	// Проверка наличия данных в файле
	if data == nil {
		for _, v := range adr {
			res, err := spr.Scan(v, *dFlag)
			if err != nil {
				log.Printf("Ошибка! %v", err)
				continue
			}
			data = append(data, res...)
		}

		data = spider.EnumId(data)

	}

	if err := inx.Add(data...); err != nil {
		log.Fatalf("ошибка! Не возможно добавить индексы. %v", err)
	}

	i := inx.Index[strings.ToLower(*sFlag)]
	if i != nil {
		for _, v := range i {
			fmt.Println(binarySearch(data, v).URL)
		}
	} else {
		fmt.Println("Links not found")
	}

	os.Exit(0)

}

//help справка по использованию утилиты
func help() {
	fmt.Printf("Использование:\n")
	fmt.Printf("  %s [аргументы]\n", filepath.Base(os.Args[0]))
	fmt.Printf("\n")
	fmt.Printf("Aргументы:\n")
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
