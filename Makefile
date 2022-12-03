day?=day1

build:
	go build -o ./dist/adventOfCode

run:
	go build -o ./dist/adventOfCode
	./dist/adventOfCode $(day)

test: 
	go test ./... -v

benchmark:
	go test ./... -v -bench=./... -benchmem