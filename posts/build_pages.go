package posts

import (
	"bytes"
	"os"

	"github.com/yuin/goldmark"
)

func Make() {
	markdownBytes, err := os.ReadFile("posts/2024/02-stop-switching-to-vscode.md")

	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(markdownBytes, &buf); err != nil {
		panic(err)
	}

	os.WriteFile("out/test.html", buf.Bytes(), 0644)
}
