package main

import (
	"embed"
	"fmt"

	"github.com/youthlin/t"
)

//go:embed language
var language embed.FS

func main() {
	t.LoadFS(language)
	fmt.Println(t.T("Hello, World"))
	t.N("I have One appale.", "I've %v apples.", 2, 2)
}
