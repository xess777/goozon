build:
	docker build -t spaces .

run:
	docker run --rm spaces run src/main.go

test:
	docker run --rm spaces test ./...
