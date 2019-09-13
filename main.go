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

	"github.com/mpuzanov/parser1c/models"
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

func importData(data string) (doc models.Document1C, err error) {
	doc = models.Document1C{}
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
