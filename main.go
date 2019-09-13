// parser1c project main.go
// go run . files/kl_to_1c.txt
// go run . files/kl_to.txt

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"os"
	"path/filepath"

	"golang.org/x/text/encoding/charmap"
)

func main() {
	//fmt.Println(os.Args)
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <file1> \n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	file := os.Args[1]

	fmt.Printf("input file: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	defer f.Close()

	readerDecoder := charmap.Windows1251.NewDecoder().Reader(f)

	if rawBytes, err := ioutil.ReadAll(readerDecoder); err != nil {
		log.Fatal(err)
	} else {
		//fmt.Println(string(rawBytes))
		importData(string(rawBytes))
	}

}

func importData(data string) (doc Document1C, err error) {
	doc = Document1C{}
	lines := strings.SplitN(data, "\n", 2)
	if strings.TrimSpace(lines[0]) != "1CClientBankExchange" {
		log.Fatal("File not 1CClientBankExchange")
		//return doc,nil
	}
	// s := ""
	// for _, line := range lines {
	// 	//s += strings.TrimSpace(line)
	// 	s += line
	// }
	// fmt.Println(s)
	fmt.Println(data)
	return doc, nil
}

// func importFile(inFile io.Reader) (err error) {
// 	reader := bufio.NewReader(inFile)
// 	eof := false
// 	for !eof {
// 		var line string
// 		line, err = reader.ReadString('\n')
// 		if err == io.EOF {
// 			err = nil  // в действительности признак io.EOF не является ошибкой
// 			eof = true // это вызовет прекращение цикла в следующей итерации
// 		} else if err != nil {
// 			return err // в случае настоящей ошибки выйти немедленно
// 		}
// 		fmt.Println(line)
// 	}
// 	return nil
// }

// func importFile2(inFile io.Reader) (err error) {

// 	dec := charmap.Windows1251.NewDecoder().Reader(inFile)

// scanner := bufio.NewScanner(readerDecoder)
// for scanner.Scan() {
// 	fmt.Println(scanner.Text())
// }

// 	return nil
// }
