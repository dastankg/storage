package main

import (
	"fmt"
	"storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("Hello World", st)
}
