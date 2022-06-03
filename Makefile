# build & test automation

build:
	go build pacman.go

test: build
	./pacman

clean:
	rm -rf pacman