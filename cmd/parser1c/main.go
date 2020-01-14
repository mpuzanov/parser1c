/*
Программа обработки банковского файла из 1С
*/

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/mpuzanov/parser1c/internal/app/parser1c"
	"golang.org/x/text/encoding/charmap"
)

var (
	format  = flag.String("format", "xlsx", "Формат вывода (по умолчанию xlsx)")
	formats = map[string]string{"json": ".json", "xlsx": ".xlsx", "csv": ".csv"}
)

func main() {
	flag.Parse()
	flag.Args()

	if len(flag.Args()) == 0 {
		fmt.Printf("usage: %s -format xlsx <file1> \n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	file := flag.Args()[0]

	fmt.Printf("файл на входе: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	defer f.Close()

	//Перекодируем из win1251
	readerDecoder := charmap.Windows1251.NewDecoder().Reader(f)
	var rawBytes []byte
	if rawBytes, err = ioutil.ReadAll(readerDecoder); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(rawBytes))

	doc, err := parser1c.ImportData(string(rawBytes))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// формируем имя файла - меняем расширение
	newFileName := strings.TrimSuffix(file, path.Ext(file)) + formats[*format]

	// Сохраняем результат в файл
	switch *format {
	case "json":
		err = doc.ToJSON(newFileName)
	case "xlsx":
		err = doc.ToExcel(newFileName)
	case "csv":
		err = doc.ToCsv(newFileName)
	default:
		fmt.Printf("Формат вывода <%s> не определён в программе", *format)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("результат: %s\n", newFileName)
}
