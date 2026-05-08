package internal

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func CreateFilesServer() *http.Server {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	http.HandleFunc("/{lang}/just", filesResponse)
	http.HandleFunc("/langages/list", langagesListResponse)

	return s
}

func langagesListResponse(w http.ResponseWriter, r *http.Request) {
	result := LangagesList()
	io.WriteString(w, fmt.Sprintf("%v", result))
}

func filesResponse(w http.ResponseWriter, r *http.Request) {
	lang := r.PathValue("lang")
	langage, err := VerifyLangageExist(lang)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
	} else {
		fmt.Printf("%s justfile served !\n", *langage)
		http.ServeFile(w, r, fmt.Sprintf("../%s/justfile", *langage))
	}
}
