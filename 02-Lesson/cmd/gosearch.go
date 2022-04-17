package main

import (
	"GoSearch/02-Lesson/pkg/crawler"
	"GoSearch/02-Lesson/pkg/crawler/spider"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	//Инициализация аргументов
	sFlag := flag.String("s", "", `поиск слова на страницах "go.dev" и "golang.org"`)
	dFlag := flag.Int("d", 0, "глубина перехода по ссылкам")
	flag.Usage = help
	flag.Parse()

	//Инициализация робота
	spr := spider.New()

	//Список адресов
	adr := []string{"https://go.dev", "https://golang.org"}

	//Результат сканированя робота
	var data []crawler.Document

	//Добавление результатов сканирования
	for _, v := range adr {
		res, err := spr.Scan(v, *dFlag)
		if err != nil {
			log.Printf("Ошибка! %v", err)
		}
		data = append(data, res...)
	}

	var urlC int //счетчик ссылок

	//Поиск слова в найденных ссылках
	for _, v := range data {
		if strings.Contains(v.URL, *sFlag) {
			fmt.Println(v.URL)
			urlC++
		}
	}

}

//help справка по использованию утилиты
func help() {
	_, _ = fmt.Fprintf(os.Stderr, "Использование:\n")
	_, _ = fmt.Fprintf(os.Stderr, "  %s [аргументы]\n", filepath.Base(os.Args[0]))
	_, _ = fmt.Fprintf(os.Stderr, "\n")

	_, _ = fmt.Fprintf(os.Stderr, "Aргументы:\n")
	flag.PrintDefaults()
}
