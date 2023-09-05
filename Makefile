.PHONY: all
all: txgen

.PHONY: txgen
txgen:
	go build -o build/txgen .


.PHONY: clean
clean:
	rm -fr ./build/*
