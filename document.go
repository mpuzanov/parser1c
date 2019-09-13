package main

// Header ...
type Header struct {
	fileVer      string // ВерсияФормата
	fileCodePage string // Кодировка
	fileDate     string // ДатаСоздания
}

// Document ...
type Document struct {
	docNum  string  // Номер
	docDate string  // Дата
	docSum  float64 // Сумма

	sourceRS   string // ПлательщикСчет
	sourceINN  string // ПлательщикИНН
	sourceName string // Плательщик

	targetRS   string // ПолучательСчет
	targetINN  string // ПолучательИНН
	targetName string // Получатель

	purpose string // НазначениеПлатежа
}

// Document1C ...
type Document1C struct {
	header    Header
	documents []Document
}
