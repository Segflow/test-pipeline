OUTPUT=bin/

install:
	mkdir bin/ || true
	go get github.com/shirou/gopsutil/mem
	go build -o ${OUTPUT}/service fact.go

start: install
	./${OUTPUT}/service

clean:
	rm -r bin

test:
	go test ./...