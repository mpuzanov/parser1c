package parser1c

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/mpuzanov/parser1c/internal/app/models"
)

var pattern string

//ImportData ...
func ImportData(data string) (doc *models.File1C, err error) {

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
