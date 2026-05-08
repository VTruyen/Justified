package main

import (
	"fmt"
	"log"

	"github.com/justified/server/internal"
)

func main() {
	s := internal.CreateFilesServer()
	fmt.Println("Server is working...!")
	log.Fatal(s.ListenAndServe())
}
