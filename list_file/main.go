package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ListFiles(dirname string, level int) {
	s := "|--"
	for i := 0; i < level; i++ {
		s = "|  " + s
	}
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatalln(err)
	}
	for _, fi := range fileInfos {
		filename := dirname + "/" + fi.Name()
		fmt.Printf("%s%s\n", s, filename)
		if fi.IsDir() {
			ListFiles(filename, level+1)
		}
	}
}
func main() {
	dirname := "/Users/topstore/go/src/gopl.io"
	ListFiles(dirname, 0)
}
