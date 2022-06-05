# build & test automation

build:
	go build ./pacman

test: build
	./pacman

clean:
	rm -rf pacman.exe