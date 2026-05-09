package main

import (
	"fmt"
	"justified/server/internal"
	"log"
)

func main() {
	s := internal.CreateFilesServer()
	fmt.Println("Server is working...!")
	log.Fatal(s.ListenAndServe())
}
