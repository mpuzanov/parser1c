package model

// File1C файл выгрузки платежей из 1С
type File1C struct {
	Header map[string]string
	Docs   []map[string]string
}
