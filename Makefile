build:
	go build app/main.go
run:
	go run app/main.go
test:
	go test ./...
benchmark:
	go test -bench=. ./... -count 5 -benchtime=10s -benchmem