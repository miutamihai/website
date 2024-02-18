gen:
	@templ generate

run:
	@go run .

css:
	@./bin/tailwindcss -i assets/input.css -o out/assets/app.css

css-ci:
	@./bin/tailwindcss-ci -i assets/input.css -o out/assets/app.css

# TODO: Change this to remove node dependency
serve:
	@npx serve -s ./out

full:
	@make gen
	@make css
	@make run

