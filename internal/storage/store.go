package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"parser1c/internal/domain/model"
	"unicode/utf8"

	"github.com/tealeg/xlsx"
)

// File1C структура с платежами из выписки 1С
type File1C model.File1C

//HeaderFile параметры файла
var HeaderFile = []string{"ВерсияФормата", "Кодировка", "ДатаСоздания"}

//HeaderDoc список полей в секции документа
var HeaderDoc = []string{"Номер", "Дата", "Сумма", "ПлательщикСчет", "ПлательщикИНН", "Плательщик", "ПолучательСчет", "ПолучательИНН", "Получатель", "НазначениеПлатежа"}

//withHeader ширина колонок
var withHeader = make(map[string]int)

// NewFile1C Создание реестра
func NewFile1C() *File1C {
	d := &File1C{}
	d.Header = make(map[string]string)
	d.Docs = make([]map[string]string, 0)
	return d
}

//CountDoc возвращает кол-во документов(платежей) в реестре
func (doc *File1C) CountDoc() int {
	return len(doc.Docs)
}

//ToCsv Сохранение в csv-формат
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
func (doc *File1C) ToExcel(fileName string) error {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()

	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		return err
	}

	headerFont := xlsx.NewFont(12, "Calibri")
	headerFont.Bold = true
	headerFont.Underline = false
	headerStyle := xlsx.NewStyle()
	headerStyle.Font = *headerFont

	dataFont := xlsx.NewFont(11, "Calibri")
	dataStyle := xlsx.NewStyle()
	dataStyle.Font = *dataFont //*xlsx.DefaultFont()

	//Зададим наименование колонок
	row = sheet.AddRow()
	for index := 0; index < len(HeaderDoc); index++ {
		cell = row.AddCell()
		cell.Value = HeaderDoc[index]
		cell.SetStyle(headerStyle)
		withHeader[HeaderDoc[index]] = utf8.RuneCountInString(HeaderDoc[index])
	}
	//fmt.Println(withHeader)
	//данные
	for index := 0; index < len(doc.Docs); index++ {
		row = sheet.AddRow()
		for j := 0; j < len(HeaderDoc); j++ {
			cell = row.AddCell()
			cell.Value = doc.Docs[index][HeaderDoc[j]]
			cell.SetStyle(dataStyle)

			if utf8.RuneCountInString(cell.Value) > withHeader[HeaderDoc[j]] {
				withHeader[HeaderDoc[j]] = utf8.RuneCountInString(cell.Value)
				//fmt.Println(cell.Value)
			}
		}
	}
	//Устанавливаем ширину колонок
	//fmt.Println(withHeader)
	for i, col := range sheet.Cols {
		col.Width = float64(withHeader[HeaderDoc[i]])
		//fmt.Println(i, HeaderDoc[i], col.Width)
	}

	err = file.Save(fileName)
	if err != nil {
		return err
	}
	return nil
}

//SaveInFile Сохраняем результат в файл
func SaveInFile(doc *File1C, newFileName string, format string) {

	var err error

	switch format {
	case "json":
		err = doc.ToJSON(newFileName)
	case "xlsx":
		err = doc.ToExcel(newFileName)
	case "csv":
		err = doc.ToCsv(newFileName)
	default:
		fmt.Printf("Формат вывода <%s> не определён в программе", format)
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("результат: %s\n", newFileName)
}
