package main

import (
	"context"
	"os"
	"website/internal/posts"
	"website/internal/views"
)

func main() {
	posts.Make()
	indexFile, err := os.Create("out/index.html")

	if err != nil {
		panic(err)
	}

	home := views.Home()

	home.Render(context.Background(), indexFile)
}
