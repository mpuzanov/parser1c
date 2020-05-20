package parser1c

import (
	"fmt"
	"regexp"

	"parser1c/internal/storage"
	"strings"
)

var pattern string

//ImportData Преобразуем текст файла с реестром в нашу структуру документа
func ImportData(data string) (doc *storage.File1C, err error) {

	doc = storage.NewFile1C()

	lines := strings.SplitN(data, "\n", 2)
	if strings.TrimSpace(lines[0]) != "1CClientBankExchange" {
		return nil, fmt.Errorf("File not 1CClientBankExchange")
	}

	for _, val := range storage.HeaderFile {
		doc.Header[val] = getValueName(val, data)
	}
	for _, docStr := range getStringDoc(data) {
		doc1 := make(map[string]string)
		for _, val := range storage.HeaderDoc {
			doc1[val] = getValueName(val, docStr)
		}
		doc.Docs = append(doc.Docs, doc1)
	}
	return doc, nil
}

//getStringDoc выделяем документы оплат в файле
func getStringDoc(data string) []string {
	var res []string
	pattern = `(СекцияДокумент=)(.+\s+)+?(КонецДокумента)`
	re := regexp.MustCompile(pattern)
	submatchall := re.FindAllString(data, -1)
	res = append(res, submatchall...)
	return res
}

//getValueName получаем значения заданных полей
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
