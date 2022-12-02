BINARY_NAME=aoc2022

build:
	go build -o ${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}
