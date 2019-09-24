/*
parser1c project main.go
go run . files/kl_to_1c.txt
go run . files/kl_to.txt

Тестирование шаблонов регулярных выражений:
https://regex101.com/
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
	"regexp"
	"strings"

	"github.com/mpuzanov/parser1c/models"
	"golang.org/x/text/encoding/charmap"
)

var (
	pattern string
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

	fmt.Printf("input file: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	defer f.Close()

	readerDecoder := charmap.Windows1251.NewDecoder().Reader(f)
	var rawBytes []byte
	if rawBytes, err = ioutil.ReadAll(readerDecoder); err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(rawBytes))
	doc, _ := importData(string(rawBytes))

	newFileName := strings.TrimSuffix(file, path.Ext(file)) + formats[*format]

	// Выводим полученный документ
	switch *format {
	case "json":
		doc.ToJSON(newFileName)
	case "xlsx":
		doc.ToExcel(newFileName)
	case "csv":
		doc.ToCsv(newFileName)
	default:
		fmt.Printf("Формат вывода <%s> не определён в программе", *format)
	}

}

func getStringDoc(data string) []string {

	var res []string

	pattern = `(СекцияДокумент=)(.+\s+)+?(КонецДокумента)`

	re := regexp.MustCompile(pattern)
	//fmt.Printf("Pattern: %v\n", re.String())
	//fmt.Println(re.MatchString(data))

	submatchall := re.FindAllString(data, -1)
	for _, element := range submatchall {
		res = append(res, element)
	}
	return res
}

func getValueName(name, data string) string {
	var res string
	pattern := fmt.Sprintf(`%s=(.+)`, name)
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(data)
	if match != nil {
		res = strings.TrimSpace(match[1])
	}
	return res
}

func getValue(pattern, data string) string {
	var res string
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(data)
	if match != nil {
		res = strings.TrimSpace(match[1])
	}
	return res
}

func importData(data string) (doc *models.File1C, err error) {

	doc = models.NewFile1C()

	lines := strings.SplitN(data, "\n", 2)
	if strings.TrimSpace(lines[0]) != "1CClientBankExchange" {
		log.Fatal("File not 1CClientBankExchange")
	}
	//fmt.Println(data)

	for _, val := range models.HeaderFile {
		doc.Header[val] = getValueName(val, data)
	}
	for _, docStr := range getStringDoc(data) {
		doc1 := make(map[string]string)
		for _, val := range models.HeaderDoc {
			doc1[val] = getValueName(val, docStr)
		}
		doc.Docs = append(doc.Docs, doc1)
	}
	return doc, nil
}
