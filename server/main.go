package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/justified/server/internal"
)

func main() {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	http.HandleFunc("/{lang}/just", filesResponse)

	fmt.Println("Server is working...!")
	log.Fatal(s.ListenAndServe())
}

func filesResponse(w http.ResponseWriter, r *http.Request) {
	lang := r.PathValue("lang")
	langage, err := internal.VerifyLangageExist(lang)
	if err != nil {
		io.WriteString(w, err.Error())
	} else {
		http.ServeFile(w, r, fmt.Sprintf("../%s/justfile", *langage))
	}
}
