gen:
	@templ generate

run:
	@go run .

css:
	@./tailwindcss -i assets/input.css -o assets/app.css
