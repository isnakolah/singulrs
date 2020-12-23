dev:
	go run httpd/main.go

test:
	go test -cover ./... -count=1