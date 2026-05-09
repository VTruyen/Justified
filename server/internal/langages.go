package internal

import "fmt"

type Langage string

const (
	Go     Langage = "go"
	Rust   Langage = "rust"
	Js     Langage = "js"
	Maven  Langage = "maven"
	Gradle Langage = "gradle"
	Python Langage = "python"
)

var langages = map[string]Langage{
	"go":     Go,
	"rust":   Rust,
	"js":     Js,
	"maven":  Maven,
	"gradle": Gradle,
	"python": Python,
}

func verifyLangageExist(langage string) (*Langage, error) {
	lang, ok := langages[langage]
	if !ok {
		return nil, fmt.Errorf("langage %s not supported", langage)
	}
	return &lang, nil
}

func LangagesList() string {
	langs := ""
	for k := range langages {
		langs = langs + "- " + k + "\n"
	}
	return langs
}
