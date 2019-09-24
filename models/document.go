package models

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tealeg/xlsx"
)

// File1C ...
type File1C struct {
	Header map[string]string
	Docs   []map[string]string
}

// NewFile1C Создание реестра
func NewFile1C() *File1C {
	d := &File1C{}
	d.Header = make(map[string]string)
	d.Docs = make([]map[string]string, 0, 0)
	return d
}

//HeaderFile ...
var HeaderFile = []string{"ВерсияФормата", "Кодировка", "ДатаСоздания"}

//HeaderDoc ...
var HeaderDoc = []string{"Номер", "Дата", "Сумма", "ПлательщикСчет", "ПлательщикИНН", "Плательщик", "ПолучательСчет", "ПолучательИНН", "Получатель", "НазначениеПлатежа"}

//ToCsv ...
func (doc *File1C) ToCsv(toFileName string) error {
	s := ""
	for index := 0; index < len(HeaderDoc); index++ {
		s += HeaderDoc[index] + ";"
	}
	s += "\n"
	// данные
	for index := 0; index < len(doc.Docs); index++ {
		for j := 0; j < len(HeaderDoc); j++ {
			s += doc.Docs[index][HeaderDoc[j]] + ";"
		}
		s += "\n"
	}
	// Записываем в файл
	f, err := os.OpenFile(toFileName, os.O_WRONLY|os.O_CREATE, 0660)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Fprintf(f, "%s\n", s)
	return nil
}

//ToJSON Возвращаем информацию в JSON формате
func (doc *File1C) ToJSON(toFileName string) error {

	jsData, _ := json.MarshalIndent(doc, "", "    ")
	//jsData, _ = json.Marshal(doc)

	// Записываем в файл
	f, err := os.OpenFile(toFileName, os.O_WRONLY|os.O_CREATE, 0660)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Fprintf(f, "%s\n", string(jsData))
	return nil
}

// ToExcel ...
func (doc *File1C) ToExcel(fileName string) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()

	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	//Зададим наименование колонок
	row = sheet.AddRow()
	for index := 0; index < len(HeaderDoc); index++ {
		cell = row.AddCell()
		cell.Value = HeaderDoc[index]
	}
	//данные
	for index := 0; index < len(doc.Docs); index++ {
		row = sheet.AddRow()
		for j := 0; j < len(HeaderDoc); j++ {
			cell = row.AddCell()
			cell.Value = doc.Docs[index][HeaderDoc[j]]
		}
	}

	err = file.Save(fileName)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
