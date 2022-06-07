# build & test automation

GHOSTNUM = 4

build:
	go build ./pacman

run: build
	./pacman --ghosts ${GHOSTNUM}

clean:
	rm -rf pacman.exe