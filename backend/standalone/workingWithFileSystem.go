package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	readFileExample("mnemonics.json")
	fmt.Println(os.Getwd())
}

func readFileExample(filename string) {
	dat, err := os.ReadFile(filename)
	check(err)
	fmt.Print(string(dat))
}

func readFilesInDir() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
