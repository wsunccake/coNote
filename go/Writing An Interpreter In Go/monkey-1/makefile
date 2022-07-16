PROJECT = monkey
EXECUTABLE = monkey.exe


.PHONY: build
build:
	make clean
	make test
	make main

main: main.go
	@echo "build..."
	go build -o ${EXECUTABLE} main.go

.PHONY: clean
clean:
	@echo "clean..."
	-test -f ${EXECUTABLE} && rm ${EXECUTABLE}

.PHONY: init
init:
	@echo "craete..."
	-@go mod init ${PROJECT}
	-@go mod tidy

.PHONY: run
run: main
	@echo "run..."
	./${EXECUTABLE}

.PHONY: test
test:
	go test monkey/lexer
