BINARY_NAME=hanif_skeleton

run-http:
	@go run main.go http

run-migration:
	@go run main.go db:migrate up