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
	go test -bench=ESbulk$$ -benchmem -cpuprofile=cpu.out
	go tool pprof json.test cpu.out

profilemem:
	go test -bench=ESbulk$$ -benchmem -memprofile=mem.out
	#go tool pprof -alloc_space json.test mem.out
	go tool pprof -alloc_objects json.test mem.out
