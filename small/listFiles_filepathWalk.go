package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	root      = "./data"
	extension = ".txt"
)

func main() {
	list_files, err := getListFiles(root, extension)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range list_files {
		fmt.Println(file)
	}
}

func getListFiles(root, extension string) (list []string, err error) {
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if extension != "*" && filepath.Ext(path) != extension {
			return nil
		}

		list = append(list, path)

		return nil
	})

	return list, err
}
