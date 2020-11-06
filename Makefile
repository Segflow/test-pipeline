OUTPUT=bin/



install:
	go get github.com/shirou/gopsutil/mem
	
build: install
	mkdir bin/ || true
	go build -o ${OUTPUT}/service fact.go

start: build
	./${OUTPUT}/service

clean:
	rm -r bin

test: install
	go test -v ./...

benchmark: install
	go test -run='^$\' -bench=. -benchmem


sleep: 
	sleep 3600