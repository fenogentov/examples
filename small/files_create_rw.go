package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Create("tst.csv")
	if err != nil {
		log.Fatalln(err)
	}
	f.WriteString("T E S T\n")
	f.Close()

	file, err := os.OpenFile("tst.csv", os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	w := csv.NewWriter(file)
	txt := []string{"#", "Tag1", "Tag2", "Tag3", "Tag4", "Tag5"}
	w.Write(txt)
	w.Flush()
}
