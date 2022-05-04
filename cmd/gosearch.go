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
		log.Fatal(err)
	}

	sFlag := flag.String("s", "", `поиск слова на страницах "go.dev" и "golang.org"`)
	dFlag := flag.Int("d", 0, "глубина перехода по ссылкам")
	pFlag := flag.String("p", exPath, "путь для сохранения дампа результата сканирования")
	flag.Usage = help
	flag.Parse()

	// Инициализация робота
	spr := spider.New()
	// Инициализация БД индексов
	inx := index.New()

	// Список адресов
	adr := []string{"https://go.dev", "https://golang.org"}

	// Результат сканированя робота
	var data []crawler.Document

	str, err := store.New(*pFlag)
	if err != nil {
		log.Println(err)
	}

	data, err = str.Data()
	if err != nil {
		log.Println(err)
	}
	// Проверка наличия данных в файле
	if data == nil {
		// Сканирование
		for _, v := range adr {
			res, err := spr.Scan(v, *dFlag)
			if err != nil {
				log.Printf("Ошибка! %v", err)
				continue
			}
			data = append(data, res...)
		}
		// Добавление идентификаторв
		data = spider.EnumId(data)

		// добавление
		err = str.Save(data)
		if err != nil {
			log.Println(err)
		}
	}

	// Добавление индексов в БД
	if err := inx.Add(data...); err != nil {
		fmt.Println(err)
		os.Exit(1)
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
