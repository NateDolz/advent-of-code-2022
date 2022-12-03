day?=day1

build:
	go build -o ./dist/adventOfCode

run:
	go build -o ./dist/adventOfCode
	./dist/adventOfCode $(day)

