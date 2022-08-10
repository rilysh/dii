all:
	make build
	make pkg

build:
	go build -ldflags "-w" dii.go

pkg:
	tar cf dii.tar ./dii
	gzip -9 dii.tar

install:
	cp dii /usr/bin