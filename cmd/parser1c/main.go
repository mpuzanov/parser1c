package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"golang.org/x/text/encoding/charmap"

	"parser1c/internal/models"
	"parser1c/internal/parser1c"
)

func main() {

	var (
		format  = flag.String("format", "xlsx", "Формат вывода (по умолчанию xlsx)")
		formats = map[string]string{"json": ".json", "xlsx": ".xlsx", "csv": ".csv"}
	)

	flag.Parse()
	flag.Args()

	if len(flag.Args()) == 0 {
		log.Printf("usage: %s -format xlsx <file1> \n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	file := flag.Args()[0]

	log.Printf("файл на входе: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//Перекодируем из win1251
	readerDecoder := charmap.Windows1251.NewDecoder().Reader(f)
	var rawBytes []byte
	if rawBytes, err = ioutil.ReadAll(readerDecoder); err != nil {
		log.Fatal(err)
	}
	//log.Println(string(rawBytes))

	doc, err := parser1c.ImportData(string(rawBytes))
	if err != nil {
		log.Fatal(err)
	}

	if doc.CountDoc() == 0 {
		//log.Println(doc)
		log.Printf("Кол-во документов в реестре %d. Выходной файл не создан!", doc.CountDoc())
		return
	}
	// формируем имя файла - меняем расширение
	newFileName := strings.TrimSuffix(file, path.Ext(file)) + formats[*format]
	models.SaveInFile(doc, newFileName, *format)

}
