package main

import (
	"context"
	"fmt"
	"os"
	"website/internal/posts"
	"website/internal/tailwind"
	"website/internal/views"

	"github.com/a-h/templ"
)

type importInput struct {
	component templ.Component
	name      string
}

func makeGoImports() []importInput {
	return []importInput{
		{
			component: views.XgoGoImport(),
			name:      "xgo",
		},
	}
}

func main() {
	if err := tailwind.Ensure(); err != nil {
		panic(err)
	}

	posts.Make()
	indexFile, err := os.Create("out/index.html")

	if err != nil {
		panic(err)
	}

	home := views.Home()

	home.Render(context.Background(), indexFile)

	for _, goImport := range makeGoImports() {
		file, err := os.Create(fmt.Sprintf("out/%s.html", goImport.name))

		if err != nil {
			panic(err)
		}

		goImport.component.Render(context.Background(), file)
	}
}
