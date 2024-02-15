package main

import (
	"bytes"
	"context"
	"github.com/yuin/goldmark"
	"os"
)

func main() {
	markdownBytes, err := os.ReadFile("test.md")

	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(markdownBytes, &buf); err != nil {
		panic(err)
	}

	os.WriteFile("test.html", buf.Bytes(), 0644)

	component := hello("John")
	component.Render(context.Background(), os.Stdout)
}
