run:
	@./tailwindcss -i static/input.css -o static/output.css
	@templ generate
	@go run cmd/main.go