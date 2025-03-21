package main

import (
	"fmt"
	"github.com/dastankg/storage/internal/storage"
	"log"
)

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("test.txt", []byte("hello world"))
	if err != nil {
		log.Fatal(err)
	}
	restoredFile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("it uploaded", restoredFile)
}
