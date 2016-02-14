.PHONY: all build test bench profile

all: build test

build:
	go build

test:
	go test
	go vet

bench:
	go test -bench=. -benchmem

profile:
	go test -bench=RTB$$ -benchmem -cpuprofile=cpu.out
	go tool pprof json.test cpu.out
