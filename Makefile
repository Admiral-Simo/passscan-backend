build:
	@go build cmd/main.go
run:
	@go run cmd/main.go
dockerize:
	@docker build -t passport_card_analyser .
	@docker run -d -p 8080:8080 passport_card_analyser
